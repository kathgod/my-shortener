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

var (
	srvAddress   *string
	bsURL        *string
	flStoragePth *string
)

func init() {
	srvAddress = flag.String("a", "localhost:8080", "SERVER_ADDRESS")
	bsURL = flag.String("b", "http://localhost:8080", "BASE_URL")
	flStoragePth = flag.String("f", "", "FILE_STORAGE_PATH")
}

func main() {
	flag.Parse()

	portNumber := MyHandler.HandParam("SERVER_ADDRESS", srvAddress)
	MyHandler.ResHandParam.BU = MyHandler.HandParam("BASE_URL", bsURL)
	MyHandler.ResHandParam.FSP = MyHandler.HandParam("FILE_STORAGE_PATH", flStoragePth)

	mapPost := make(map[string]string)
	mapGet := make(map[string]string)

	resP := MyHandler.PostFunc(mapPost, mapGet)
	resG := MyHandler.GetFunc(mapPost, mapGet)
	resNam := MyHandler.NotAllowedMethodFunc()
	resPAS := MyHandler.PostFuncAPIShorten(mapPost, mapGet)
	resGAUU := MyHandler.GetFuncApiUserUrls(mapPost, mapGet)

	rand.Seed(time.Now().UnixNano())

	rtr := chi.NewRouter()
	rtr.Get("/{id}", resG)
	rtr.Post("/", resP)
	rtr.Post("/api/shorten", resPAS)
	rtr.MethodNotAllowed(resNam)
	rtr.Get("/api/user/urls", resGAUU)

	err2 := http.ListenAndServe(portNumber, rtr)
	if err2 != nil {
		log.Println(srError)
	}

}
