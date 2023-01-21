package main

import (
	"flag"
	"github.com/go-chi/chi/v5"
	"log"
	"math/rand"
	"net/http"
	"time"
	MyHandler "urlshortener/internal/app"
)

const srError = "Server Error"

// MyTest функция для проверки тестирования
func myTest(n int) int {
	return n * 2
}

func main() {
	flag.Parse()
	portNumber := MyHandler.HandParam("SERVER_ADDRESS")
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
