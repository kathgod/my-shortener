package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	MyHandler "urlshortener/internal/app"
)

const srError = "Server Error"

// const portNumber = ":8080"

// MyTest функция для проверки тестирования
func myTest(n int) int {
	return n * 2
}

func main() {
	portNumber := os.Getenv("SERVER_ADDRESS")
	if portNumber == "" {
		portNumber = "localhost:8080"
	}
	mapPost := make(map[string]string)
	mapGet := make(map[string]string)
	resP := MyHandler.PostFunc(mapPost, mapGet)
	resG := MyHandler.GetFunc(mapPost, mapGet)
	resNam := MyHandler.NotAllowedMethodFunc()
	resPAS := MyHandler.PostFuncAPIShorten(mapPost, mapGet)
	rand.Seed(time.Now().UnixNano())
	rtr := chi.NewRouter()
	rtr.Get("/{id}", resG)
	rtr.Post("/", resP)
	rtr.Post("/api/shorten", resPAS)
	rtr.MethodNotAllowed(resNam)
	//fmt.Printf("Starting application on port %v\n", portNumber)
	err2 := http.ListenAndServe(portNumber, rtr)
	if err2 != nil {
		log.Println(srError)
	}

}
