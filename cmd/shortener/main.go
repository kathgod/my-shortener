package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"math/rand"
	"net/http"
	"time"
	MyHandler "urlshortener/internal/app"
	//"github.com/go-chi/chi/v5/middleware"
)

const portNumber = ":8080"

func MyTest(n int) int {
	return n * 2
}

func main() {
	map_post := make(map[string]string)
	map_get := make(map[string]string)

	res_p := MyHandler.PostFunc(map_post, map_get)
	res_g := MyHandler.GetFunc(map_post, map_get)
	res_NAM := MyHandler.NAMfunc()

	rand.Seed(time.Now().UnixNano())

	//http.HandleFunc("/", res_p)

	rtr := chi.NewRouter()
	//rtr.Use(middleware.Logger)
	//rtr.Use(middleware.Recoverer)

	rtr.Get("/{id}", res_g)

	rtr.Post("/", res_p)

	rtr.MethodNotAllowed(res_NAM)

	fmt.Printf("Starting application on port %v\n", portNumber)

	http.ListenAndServe(portNumber, rtr)
	//http.ListenAndServe(portNumber, nil)

}
