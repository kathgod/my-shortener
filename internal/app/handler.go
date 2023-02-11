package handler

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/hmac"
	cr "crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Константы ошибок
const (
	postBodyError        = "Bad Post request body"
	notAllowMethodError  = "Not Allow method Error "
	closeFileError       = "Close File Error"
	writeFileError       = "Write into the File"
	seekError            = "Seek Error"
	openFileError        = "Open File Error"
	compressError        = "Compress file"
	coockieByteReadError = "Coockie Byte Read Error"
	baseurl              = "http://localhost:8080/"
	errorCreatingTable   = "Error when creating table"
	errorPrepareContext  = "Prepare context Error"
	errInsert            = "Error when inserting row into table"
	findingRowAffected   = "Error when finding rows affected"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Функция для формирования случайной поледовательности
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// GetFunc Обработчик для Get запросов
func GetFunc(_, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fileStoragePath := ResHandParam.FSP
		storageFile, fileError := os.OpenFile(fileStoragePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		if fileError != nil {
			log.Println(openFileError)
		}
		defer func(storageFile *os.File) {
			err := storageFile.Close()
			if err != nil {
				log.Println(closeFileError)
			}
		}(storageFile)
		if fileStoragePath != "" {
			count := 0
			for range handMapGet {
				count++
			}
			if count == 0 {
				mokMap := map[string]string{}
				recovery(mokMap, handMapGet, storageFile)
			}
		}

		urlGet := r.URL.Path
		out := strings.Replace(urlGet, "/", "", -1)
		if handMapGet[out] != "" {
			w.Header().Set("Location", handMapGet[out])
			w.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

// PostFunc Обработчик Post запросов
func PostFunc(handMapPost map[string]string, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bp, err := decompress(io.ReadAll(r.Body))
		if err != nil {
			log.Println(postBodyError)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			cck, errCck := r.Cookie("userId")
			cckValue := ""
			if errCck != nil {
				cChVar := coockieCheck(w, r)
				cckValue = cChVar
			} else {
				cckValue = cck.Value
			}
			log.Println("cckValue in PostFunc:", cckValue)
			resultPost, sqlError := shortPostFunc(handMapPost, handMapGet, bp, cckValue)
			bResultPost := []byte(resultPost)
			if r.Header.Get("Content-Encoding ") == "gzip" {
				bResultPost, err = compress([]byte(resultPost))
				if err != nil {
					log.Println(compressError)
				}
				w.Header().Set("Accept-Encoding", "gzip")
			}
			if sqlError != 0 {
				w.WriteHeader(http.StatusCreated)
			} else {
				w.WriteHeader(http.StatusConflict)
			}
			_, err := w.Write(bResultPost)
			if err != nil {
				http.Error(w, "Post request error", http.StatusBadRequest)

			}
		}
	}
}

// Функуция сокращения URL для PostFunc
func shortPostFunc(handMapPost map[string]string, handMapGet map[string]string, bp []byte, cckValue string) (string, int64) {
	fileStoragePath := ResHandParam.FSP
	storageFile, fileError := os.OpenFile(fileStoragePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if fileError != nil {
		log.Println(openFileError)
	}
	defer func(storageFile *os.File) {
		err := storageFile.Close()
		if err != nil {
			log.Println(closeFileError)
		}
	}(storageFile)
	if fileStoragePath != "" {
		count := 0
		for range handMapPost {
			count++
		}
		if count == 0 {
			recovery(handMapPost, handMapGet, storageFile)
		}
	}
	baseURL := ResHandParam.BU
	rndRes := randSeq(6) + cckValue
	for {
		if handMapGet[string(bp)] != "" {
			rndRes = randSeq(6) + cckValue
		} else {
			break
		}
	}

	resultPost := baseURL + rndRes

	var sqlError int64 = -1
	if ResHandParam.DBD != "" {
		sqlError = AddRecordInTable(ResCreateSQLTable, resultPost, string(bp), cckValue)
		log.Println(sqlError)
	}
	if sqlError != 0 {
		handMapPost[string(bp)] = rndRes
		handMapGet[rndRes] = string(bp)
		addToFile := string(bp) + "@" + rndRes + "\n"
		if fileStoragePath != "" {
			_, err2 := storageFile.Write([]byte(addToFile))
			if err2 != nil {
				log.Println(writeFileError)
			}
		}
	} else {
		buff := handMapPost[string(bp)]
		resultPost = baseURL + buff
	}

	return resultPost, sqlError
}

// NotAllowedMethodFunc Обработчик для незаданных методов
func NotAllowedMethodFunc() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(notAllowMethodError)
		w.WriteHeader(http.StatusBadRequest)
	}
}

// URLLongAndShort Структура для джейсон объектов
type URLLongAndShort struct {
	OriginalURL string `json:"url,omitempty"`
	ShortURL    string `json:"result,omitempty"`
}

// PostFuncAPIShorten бработчик Post запросов для эндпоинта api/shorten/
func PostFuncAPIShorten(handMapPost map[string]string, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rawBsp, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(postBodyError)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			shURLByteFormat, err0 := shortPostFuncAPIShorten(handMapPost, handMapGet, rawBsp)
			w.Header().Set("Content-Type", "application/json")
			if err0 != 0 {
				w.WriteHeader(http.StatusCreated)
			} else {
				w.WriteHeader(http.StatusConflict)
			}

			_, err1 := w.Write(shURLByteFormat)
			if err1 != nil {
				http.Error(w, "Post request error", http.StatusBadRequest)
			}

		}

	}
}

// Функция сокращения URL для PostFuncAPIShorten
func shortPostFuncAPIShorten(handMapPost map[string]string, handMapGet map[string]string, rawBsp []byte) ([]byte, int64) {
	fileStoragePath := ResHandParam.FSP
	storageFile, fileError := os.OpenFile(fileStoragePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if fileError != nil {
		log.Println(openFileError)
	}
	defer func(storageFile *os.File) {
		err := storageFile.Close()
		if err != nil {
			log.Println(closeFileError)
		}
	}(storageFile)
	if fileStoragePath != "" {
		count := 0
		for range handMapPost {
			count++
		}
		if count == 0 {
			recovery(handMapPost, handMapGet, storageFile)
		}
	}

	baseURL := ResHandParam.BU
	urlStruct := URLLongAndShort{}
	if err := json.Unmarshal(rawBsp, &urlStruct); err != nil {
		log.Println(postBodyError)
	}
	rndRes := randSeq(6)
	for {
		if handMapGet[rndRes] != "" {
			rndRes = randSeq(6)
		} else {
			break
		}
	}

	urlStruct.ShortURL = baseURL + rndRes
	var sqlErr int64 = -1
	if ResHandParam.DBD != "" {
		sqlErr = AddRecordInTable(ResCreateSQLTable, urlStruct.ShortURL, urlStruct.OriginalURL, "default")
		log.Println(sqlErr)
	}
	var shURLByteFormat []byte
	if sqlErr != 0 {
		handMapPost[urlStruct.OriginalURL] = rndRes
		handMapGet[rndRes] = urlStruct.OriginalURL

		addToFile := urlStruct.OriginalURL + "@" + rndRes + "\n"
		if fileStoragePath != "" {
			_, err := storageFile.Write([]byte(addToFile))
			if err != nil {
				log.Println(writeFileError)
			}
		}
	} else {
		buff := handMapPost[urlStruct.OriginalURL]
		urlStruct.ShortURL = baseURL + buff
	}
	urlStruct.OriginalURL = ""
	shURLByteFormat, _ = json.Marshal(urlStruct)

	return shURLByteFormat, sqlErr
}

// Функия востановления данных из файла
func recovery(handMapPost map[string]string, handMapGet map[string]string, file *os.File) {
	_, err := file.Seek(0, 0)
	if err != nil {
		log.Println(seekError)
	}
	mReader := bufio.NewReader(file)
	for {
		data1, err1 := mReader.ReadBytes('@')
		data2, err2 := mReader.ReadBytes('\n')
		if err1 != nil || err2 != nil {
			break
		}
		handMapPost[strings.Replace(string(data1), "@", "", -1)] = strings.Replace(string(data2), "\n", "", -1)
		handMapGet[strings.Replace(string(data2), "\n", "", -1)] = strings.Replace(string(data1), "@", "", -1)
	}

}

// HandParam Функция обработки флагов
func HandParam(name string, flg *string) string {
	res := ""
	globEnv := os.Getenv(name)
	if globEnv != "" {
		res = globEnv
	} else {
		res = *flg
	}
	switch name {
	case "SERVER_ADDRESS":
	case "BASE_URL":
		res = res + "/"
	case "FILE_STORAGE_PATH":
	case "DATABASE_DSN":
	}
	return res
}

// ResHandParam Структура для предобработки флагов и переменных
var ResHandParam struct {
	BU  string
	FSP string
	DBD string
}

// Функция декомпресии тела запроса
func decompress(data []byte, err0 error) ([]byte, error) {
	if err0 != nil {
		return nil, fmt.Errorf("error 0 %v", err0)
	}

	r, err1 := gzip.NewReader(bytes.NewReader(data))
	if err1 != nil {
		return data, nil
	}
	defer r.Close()

	var b bytes.Buffer

	_, err := b.ReadFrom(r)
	if err != nil {
		return data, nil
	}

	return b.Bytes(), nil
}

// Функция компресии тела ответа
func compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err := w.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed write data to compress temporary buffer: %v", err)
	}
	err0 := w.Close()
	if err0 != nil {
		return nil, fmt.Errorf("%v", err0)
	}
	return b.Bytes(), nil
}

// Структура для мапы сохранений куки
type idKey struct {
	id  string
	key string
}

// Мапа для сохранения куки
var resIDKey = map[string]idKey{"0": {"0", "0"}}

// Функция проверки наличия и подписи куки
func coockieCheck(w http.ResponseWriter, r *http.Request) string {
	cck, err := r.Cookie("userId")
	if err != nil {
		log.Println("Error1 Coockie check", err)
		resCCh := makeNewCoockie(w)
		return resCCh
	} else {
		rik := resIDKey[cck.Value]
		id := []byte(rik.id)
		key := []byte(rik.key)
		h := hmac.New(sha256.New, key)
		h.Write(id)
		sgnIDKey := h.Sum(nil)
		if hex.EncodeToString(sgnIDKey) != cck.Value {
			resCCh := makeNewCoockie(w)
			return resCCh
		}
	}
	return cck.Value
}

// Функция для создания новых куки при провале проверки
func makeNewCoockie(w http.ResponseWriter) string {
	id := make([]byte, 16)
	key := make([]byte, 16)
	_, err1 := cr.Read(id)
	_, err2 := cr.Read(key)

	if err1 != nil || err2 != nil {
		log.Println(coockieByteReadError)
	}
	h := hmac.New(sha256.New, key)
	h.Write(id)
	sgnIDKey := h.Sum(nil)
	cck := &http.Cookie{
		Name:  "userId",
		Value: hex.EncodeToString(sgnIDKey),
	}
	http.SetCookie(w, cck)
	resIDKey[hex.EncodeToString(sgnIDKey)] = idKey{hex.EncodeToString(id), hex.EncodeToString(key)}
	return hex.EncodeToString(sgnIDKey)
}

// OrShURL Структура для Json массива, необходимого для вывода по запросу GetFuncApiUserUrls
type OrShURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// GetFuncAPIUserUrls функция возвращает объект json-array, со всеми длинными и короткими URL которые создал юзер
func GetFuncAPIUserUrls(_, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cck, err := r.Cookie("userId")
		cckValue := ""
		if err != nil {
			cChvar := coockieCheck(w, r)
			cckValue = cChvar
		} else {
			cckValue = cck.Value
		}
		log.Println("cChVar in GetApiFunc:", cckValue)
		bm := make(map[string]string)
		for k, v := range handMapGet {
			log.Println(k[6:], v)
			if k[6:] == cckValue {
				bm[k] = v
			}
		}
		log.Println(len(bm))
		if len(bm) == 0 {
			w.WriteHeader(http.StatusNoContent)
		} else {
			var buff1 OrShURL
			buff2 := make([]OrShURL, len(bm))
			i := 0
			for k := range bm {

				buff1 = OrShURL{ShortURL: baseurl + k, OriginalURL: bm[k]}
				buff2[i] = buff1
				i++

			}
			buff3, _ := json.Marshal(buff2)
			fmt.Println(string(buff3))

			w.Header().Set("Content-Type", "application/json")
			w.Write(buff3)

		}
	}
}

// GetFuncPing Функция пинга базы данных
func GetFuncPing(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("In func", ResHandParam.DBD)
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := db.PingContext(ctx); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

// CreateSQLTable Функция создания SQL таблиц
func CreateSQLTable(db *sql.DB) *sql.DB {
	query := `CREATE TABLE IF NOT EXISTS idshortlongurl(shorturl text , longurl text primary key, userid text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Println(errorCreatingTable)
		log.Println(err)
	}
	rows, err2 := res.RowsAffected()
	if err2 != nil {
		log.Println(findingRowAffected)
	}
	log.Printf("%d rows created CreateSQLTable", rows)
	return db
}

// ResCreateSQLTable переменная для записи результата функции CreateSQLTable
var ResCreateSQLTable *sql.DB

// AddRecordInTable Функция записи в SQL таблицу
func AddRecordInTable(db *sql.DB, shortURL string, longURL string, userID string) int64 {
	query := `INSERT INTO idshortlongurl(shorturl, longurl, userid) VALUES ($1, $2, $3) ON CONFLICT (longurl) DO NOTHING`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelfunc()
	stmt, err0 := db.PrepareContext(ctx, query)
	if err0 != nil {
		log.Println(errorPrepareContext)
		log.Println(err0)
	}
	defer stmt.Close()
	res, err1 := stmt.ExecContext(ctx, shortURL, longURL, userID)
	if err1 != nil {
		log.Println(errInsert)
	}
	rows, err2 := res.RowsAffected()
	if err2 != nil {
		log.Println(findingRowAffected)
	}
	log.Printf("%d rows created AddRecordInTable", rows)
	return rows
}

type LngShrtCrltnID struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
	ShortURL      string `json:"short_url"`
}

func PostFuncAPIShortenBatch(handMapPost map[string]string, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bp, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(postBodyError)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			cck, errCck := r.Cookie("userId")
			cckValue := ""
			if errCck != nil {
				cChVar := coockieCheck(w, r)
				cckValue = cChVar
			} else {
				cckValue = cck.Value
			}
			log.Println("cckValue in PostFunc:", cckValue)
			resultPostAPIShortenBatch := shortPostAPIShortenBatch(handMapPost, handMapGet, bp)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(resultPostAPIShortenBatch)
		}
	}
}

func shortPostAPIShortenBatch(handMapPost map[string]string, handMapGet map[string]string, bp []byte) []byte {
	var postAPIShortenBatchMass []LngShrtCrltnID
	if err := json.Unmarshal(bp, &postAPIShortenBatchMass); err != nil {
		log.Println(postBodyError)
	}
	for i := 0; i < len(postAPIShortenBatchMass); i++ {
		buff := randSeq(6)
		handMapPost[postAPIShortenBatchMass[i].OriginalURL] = buff
		handMapGet[buff] = postAPIShortenBatchMass[i].OriginalURL
		postAPIShortenBatchMass[i].ShortURL = baseurl + buff
		postAPIShortenBatchMass[i].OriginalURL = ""
	}
	buff, _ := json.Marshal(postAPIShortenBatchMass)
	return buff
}
