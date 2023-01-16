package handler

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
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
			resultPost := "http://localhost:8080/" + rndRes
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

type UrlLongAndShort struct {
	OriginalUrl string `json:"url,omitempty"`
	ShortUrl    string `json:"result,omitempty"`
}

func PostFuncApiShorten(handMapPost map[string]string, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		urlStruct := UrlLongAndShort{}
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
				if handMapGet[urlStruct.OriginalUrl] != "" {
					rndRes = randSeq(6)
				} else {
					break
				}
			}
			handMapPost[urlStruct.OriginalUrl] = rndRes
			handMapGet[rndRes] = urlStruct.OriginalUrl
			urlStruct.OriginalUrl = ""
			urlStruct.ShortUrl = rndRes
			shUrlByteFormat, _ := json.Marshal(urlStruct)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			//w.Header().Set("Content-Type", "application/json")
			_, err := w.Write(shUrlByteFormat)
			if err != nil {
				http.Error(w, "Post request error", http.StatusBadRequest)
			}

		}

	}
}
