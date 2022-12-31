package handler

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string { //Функция для формирования случайной поледовательности
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetFunc(_, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) { //Обработчик для Get запросов

	return func(w http.ResponseWriter, r *http.Request) {
		urlGet := r.URL.Path
		out := strings.Replace(urlGet, "/", "", -1)
		if handMapGet[out] != "" {
			w.Header().Set("Location", handMapGet[out])
			w.WriteHeader(307)
		} else {
			w.WriteHeader(400)
		}
	}
}

func PostFunc(handMapPost map[string]string, handMapGet map[string]string) func(w http.ResponseWriter, r *http.Request) { //Обработчик Post запросов

	return func(w http.ResponseWriter, r *http.Request) {

		bp, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error 11")
			w.WriteHeader(400)
		} else {
			rndRes := RandSeq(6)
			handMapPost[string(bp)] = rndRes
			handMapGet[rndRes] = string(bp)
			resultPost := "http://localhost:8080/" + rndRes
			w.WriteHeader(201)
			_, err := w.Write([]byte(resultPost))
			if err != nil {
				os.Exit(50)
			}
		}
	}
}

func NotAllowedMethodFunc() func(w http.ResponseWriter, r *http.Request) { //Обработчик для незаданных методов

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Error method")
		w.WriteHeader(400)
	}
}
