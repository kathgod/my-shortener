package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"math/rand"
	"net/http"
	"time"
	MyHandler "urlshortener/internal/app"
)

const srError = "Server Error"
const portNumber = ":8080"

// MyTest функция для проверки тестирования
func myTest(n int) int {
	return n * 2
}

func main() {
	mapPost := make(map[string]string)
	mapGet := make(map[string]string)
	resP := MyHandler.PostFunc(mapPost, mapGet)
	resG := MyHandler.GetFunc(mapPost, mapGet)
	resNam := MyHandler.NotAllowedMethodFunc()
	resPAS := MyHandler.PostFuncApiShorten(mapPost, mapGet)
	rand.Seed(time.Now().UnixNano())
	rtr := chi.NewRouter()
	rtr.Get("/{id}", resG)
	rtr.Post("/", resP)
	rtr.Post("/api/shorten", resPAS)
	rtr.MethodNotAllowed(resNam)
	fmt.Printf("Starting application on port %v\n", portNumber)
	err := http.ListenAndServe(portNumber, rtr)
	if err != nil {
		log.Println(srError)
	}

}
