package handler

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const postBodyError = "Bad Post request body"
const notAllowMethodError = "Not Allow method Error "

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
		fileStoragePath := os.Getenv("FILE_STORAGE_PATH")
		storageFile, fileError := os.OpenFile(fileStoragePath, os.O_RDWR|os.O_APPEND, 0777)
		defer storageFile.Close()
		if fileError == nil {
			count := 0
			for range handMapPost {
				count++
			}
			if count == 0 {
				Recovery(handMapPost, handMapGet, storageFile)
			}
		}
		baseURL := os.Getenv("BASE_URL")
		if baseURL == "" {
			baseURL = "http://localhost:8080/"
		} else {
			baseURL = baseURL + "/"
		}
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
			storageFile.Write([]byte(addToFile))
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
		fileStoragePath := os.Getenv("FILE_STORAGE_PATH")
		storageFile, fileError := os.OpenFile(fileStoragePath, os.O_RDWR|os.O_APPEND, 0777)
		defer storageFile.Close()
		if fileError == nil {
			count := 0
			for range handMapPost {
				count++
			}
			if count == 0 {
				Recovery(handMapPost, handMapGet, storageFile)
			}
		}

		baseURL := os.Getenv("BASE_URL")
		if baseURL == "" {
			baseURL = "http://localhost:8080/"
		} else {
			baseURL = baseURL + "/"
		}
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
			storageFile.Write([]byte(addToFile))
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

func Recovery(handMapPost map[string]string, handMapGet map[string]string, file *os.File) {

	file.Seek(0, 0)
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
