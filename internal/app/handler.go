package handler

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const postBodyError = "Bad Post request body"
const notAllowMethodError = "Not Allow method Error "
const closeFileError = "Close File Error"
const writeFileError = "Write into the File"
const seekError = "Seek Error"

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
		fileStoragePath := HandParam("FILE_STORAGE_PATH")
		storageFile, fileError := os.OpenFile(fileStoragePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		if fileError == nil {
			count := 0
			for range handMapGet {
				count++
			}
			if count == 0 {
				mokMap := map[string]string{}
				recovery(mokMap, handMapGet, storageFile)
			}
		}
		defer func(storageFile *os.File) {
			err := storageFile.Close()
			if err != nil {
				log.Println(closeFileError)
			}
		}(storageFile)
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
		fileStoragePath := HandParam("FILE_STORAGE_PATH")
		storageFile, fileError := os.OpenFile(fileStoragePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		if fileError == nil {
			count := 0
			for range handMapPost {
				count++
			}
			if count == 0 {
				recovery(handMapPost, handMapGet, storageFile)
			}
		}
		defer func(storageFile *os.File) {
			err := storageFile.Close()
			if err != nil {
				log.Println(closeFileError)
			}
		}(storageFile)
		baseURL := HandParam("BASE_URL")
		bp, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(postBodyError)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			rndRes := randSeq(6)
			for {
				if handMapGet[string(bp)] != "" {
					rndRes = randSeq(6)
				} else {
					break
				}
			}
			handMapPost[string(bp)] = rndRes
			handMapGet[rndRes] = string(bp)
			addToFile := string(bp) + "@" + rndRes + "\n"
			if fileStoragePath != "" {
				_, err2 := storageFile.Write([]byte(addToFile))
				if err2 != nil {
					log.Println(writeFileError)
				}
			}
			resultPost := baseURL + rndRes
			w.WriteHeader(http.StatusCreated)
			_, err := w.Write([]byte(resultPost))
			if err != nil {
				http.Error(w, "Post request error", http.StatusBadRequest)
				//os.Exit(50)
			}
		}
	}
}

// NotAllowedMethodFunc Обработчик для незаданных методов
func NotAllowedMethodFunc() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(notAllowMethodError)
		w.WriteHeader(http.StatusBadRequest)
	}
}

type URLLongAndShort struct {
	OriginalURL string `json:"url,omitempty"`
	ShortURL    string `json:"result,omitempty"`
}

// PostFuncAPIShorten бработчик Post запросов для эндпоинта api/shorten/
func PostFuncAPIShorten(handMapPost map[string]string, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fileStoragePath := HandParam("FILE_STORAGE_PATH")
		storageFile, fileError := os.OpenFile(fileStoragePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		if fileError == nil {
			count := 0
			for range handMapPost {
				count++
			}
			if count == 0 {
				recovery(handMapPost, handMapGet, storageFile)
			}
		}
		defer func(storageFile *os.File) {
			err := storageFile.Close()
			if err != nil {
				log.Println(closeFileError)
			}
		}(storageFile)
		baseURL := HandParam("BASE_URL")
		urlStruct := URLLongAndShort{}
		rawBsp, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(postBodyError)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			if err := json.Unmarshal([]byte(rawBsp), &urlStruct); err != nil {
				log.Println(postBodyError)
				w.WriteHeader(http.StatusBadRequest)
			}
			rndRes := randSeq(6)
			for {
				if handMapGet[rndRes] != "" {
					rndRes = randSeq(6)
				} else {
					break
				}
			}

			handMapPost[urlStruct.OriginalURL] = rndRes
			handMapGet[rndRes] = urlStruct.OriginalURL
			addToFile := urlStruct.OriginalURL + "@" + rndRes + "\n"
			if fileStoragePath != "" {
				_, err := storageFile.Write([]byte(addToFile))
				if err != nil {
					log.Println(writeFileError)
				}
			}
			urlStruct.OriginalURL = ""
			urlStruct.ShortURL = baseURL + rndRes
			shURLByteFormat, _ := json.Marshal(urlStruct)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, err := w.Write(shURLByteFormat)
			if err != nil {
				http.Error(w, "Post request error", http.StatusBadRequest)
			}

		}

	}
}

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

func HandParam(n string) string {
	res := ""
	globEnv := os.Getenv(n)
	if globEnv != "" {
		res = globEnv
	} else {
		srvAddress := flag.String("a", "localhost:8080", "SERVER_ADDRESS")
		bsURL := flag.String("b", "http://localhost:8080", "BASE_URL")
		flStoragePth := flag.String("f", "", "FILE_STORAGE_PATH")
		flag.Parse()
		switch n {
		case "SERVER_ADDRESS":
			res = *srvAddress
		case "BASE_URL":
			res = *bsURL + "/"
		case "FILE_STORAGE_PATH":
			res = *flStoragePth
		}
	}
	return res
}
