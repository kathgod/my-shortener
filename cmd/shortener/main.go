package main

import (
	"database/sql"
	"flag"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"time"
	MyHandler "urlshortener/internal/app"
)

const (
	srError     = "Server Error"
	dbOpenError = "Open DataBase Error"
)

var (
	srvAddress   *string
	bsURL        *string
	flStoragePth *string
	datadbaseDsn *string
)

func init() {
	srvAddress = flag.String("a", "localhost:8080", "SERVER_ADDRESS")
	bsURL = flag.String("b", "http://localhost:8080", "BASE_URL")
	flStoragePth = flag.String("f", "", "FILE_STORAGE_PATH")
	datadbaseDsn = flag.String("d", "", "DATABASE_DSN")
}

func main() {
	flag.Parse()

	portNumber := MyHandler.HandParam("SERVER_ADDRESS", srvAddress)
	MyHandler.ResHandParam.BU = MyHandler.HandParam("BASE_URL", bsURL)
	MyHandler.ResHandParam.FSP = MyHandler.HandParam("FILE_STORAGE_PATH", flStoragePth)
	MyHandler.ResHandParam.DBD = MyHandler.HandParam("DATABASE_DSN", datadbaseDsn)

	mapPost := make(map[string]string)
	mapGet := make(map[string]string)

	resP := MyHandler.PostFunc(mapPost, mapGet)
	resG := MyHandler.GetFunc(mapPost, mapGet)
	resNam := MyHandler.NotAllowedMethodFunc()
	resPAS := MyHandler.PostFuncAPIShorten(mapPost, mapGet)
	resGAUU := MyHandler.GetFuncApiUserUrls(mapPost, mapGet)
	resPFASB := MyHandler.PostFuncApiShortenBatch(mapPost, mapGet)

	rand.Seed(time.Now().UnixNano())

	rtr := chi.NewRouter()
	rtr.Get("/{id}", resG)
	rtr.Post("/", resP)
	rtr.Post("/api/shorten", resPAS)
	rtr.MethodNotAllowed(resNam)
	rtr.Get("/api/user/urls", resGAUU)
	rtr.Post("/api/shorten/batch", resPFASB)

	if MyHandler.ResHandParam.DBD != "" {
		db, errDB := sql.Open("postgres", MyHandler.ResHandParam.DBD)
		defer db.Close()
		if errDB != nil {
			log.Println(dbOpenError)
		}
		resGP := MyHandler.GetFuncPing(db)
		rtr.Get("/ping", resGP)

		MyHandler.ResCreateSQLTable = MyHandler.CreateSQLTable(db)
		log.Println(reflect.TypeOf(MyHandler.ResCreateSQLTable))
	}

	err2 := http.ListenAndServe(portNumber, rtr)
	if err2 != nil {
		log.Println(srError)
	}

}
