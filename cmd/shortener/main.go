package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"math/rand"
	"net/http"
	"os"
	"time"
	MyHandler "urlshortener/internal/app"
	//"github.com/go-chi/chi/v5/middleware"
)

const portNumber = ":8080"

func MyTest(n int) int { //функция для проверки тесирования
	return n * 2
}

func main() {
	mapPost := make(map[string]string)
	mapGet := make(map[string]string)
	resP := MyHandler.PostFunc(mapPost, mapGet)
	resG := MyHandler.GetFunc(mapPost, mapGet)
	resNam := MyHandler.NotAllowedMethodFunc()
	rand.Seed(time.Now().UnixNano())
	//http.HandleFunc("/", res_p)
	rtr := chi.NewRouter()
	//rtr.Use(middleware.Logger)
	//rtr.Use(middleware.Recovered)
	rtr.Get("/{id}", resG)
	rtr.Post("/", resP)
	rtr.MethodNotAllowed(resNam)
	fmt.Printf("Starting application on port %v\n", portNumber)
	err := http.ListenAndServe(portNumber, rtr)
	if err != nil {
		os.Exit(100)
	}

	//http.ListenAndServe(portNumber, nil)

}
