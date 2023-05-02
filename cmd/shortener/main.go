package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"reflect"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

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
	enableHTTPS  *string
	configFile   *string
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func init() {
	srvAddress = flag.String("a", "localhost:8080", "copy in SERVER_ADDRESS param")
	bsURL = flag.String("b", "http://localhost:8080", "copy in BASE_URL param")
	flStoragePth = flag.String("f", "", "copy in FILE_STORAGE_PATH param")
	datadbaseDsn = flag.String("d", "", "copy in DATABASE_DSN param")
	enableHTTPS = flag.String("s", "false", "copy in ENABLE_HTTPS param")
	configFile = flag.String("c", "", "copy in CONFIG param")
}

func main() {
	flag.Parse()

	fmt.Printf("Build version: %s\n", buildVersion)
	fmt.Printf("Build date: %s\n", buildDate)
	fmt.Printf("Build commit: %s\n", buildCommit)

	MyHandler.ResHandParam.PortNumber = MyHandler.HandParam("SERVER_ADDRESS", srvAddress)
	MyHandler.ResHandParam.BaseURL = MyHandler.HandParam("BASE_URL", bsURL)
	MyHandler.ResHandParam.FileStoragePath = MyHandler.HandParam("FILE_STORAGE_PATH", flStoragePth)
	MyHandler.ResHandParam.DataBaseDSN = MyHandler.HandParam("DATABASE_DSN", datadbaseDsn)

	enableHTTPSBuff := MyHandler.HandParam("ENABLE_HTTPS", enableHTTPS)
	var err error
	MyHandler.ResHandParam.EnableHTTPS, err = strconv.ParseBool(enableHTTPSBuff)
	if err != nil {
		MyHandler.ResHandParam.EnableHTTPS = false
	}

	config := MyHandler.HandParam("CONFIG", configFile)
	MyHandler.HandConfigParam(config)

	log.Println(MyHandler.ResHandParam)

	mapPost := make(map[string]string)
	mapGet := make(map[string]string)

	resP := MyHandler.PostFunc(mapPost, mapGet)
	resG := MyHandler.GetFunc(mapGet)
	resPAS := MyHandler.PostFuncAPIShorten(mapPost, mapGet)
	resGAUU := MyHandler.GetFuncAPIUserUrls(mapGet)
	resPFASB := MyHandler.PostFuncAPIShortenBatch(mapPost, mapGet)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	rtr := chi.NewRouter()
	rtr.Get("/{id}", resG)
	rtr.Post("/", resP)
	rtr.Post("/api/shorten", resPAS)
	rtr.Get("/api/user/urls", resGAUU)
	rtr.Post("/api/shorten/batch", resPFASB)

	rtr.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	rtr.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	rtr.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	rtr.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	rtr.Handle("/debug/pprof/block", pprof.Handler("block"))
	rtr.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))

	rtr.HandleFunc("/debug/pprof/", pprof.Index)
	rtr.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	rtr.HandleFunc("/debug/pprof/profile", pprof.Profile)
	rtr.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	rtr.HandleFunc("/debug/pprof/trace", pprof.Trace)

	if MyHandler.ResHandParam.DataBaseDSN != "" {
		db, errDB := sql.Open("postgres", MyHandler.ResHandParam.DataBaseDSN)
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Println(err)
			}
		}(db)
		if errDB != nil {
			log.Println(dbOpenError)
		}
		resGP := MyHandler.GetFuncPing(db)
		resDAUU := MyHandler.DeleteFuncAPIUserURLs(mapPost, mapGet, db, MyHandler.ResHandParam.DataBaseDSN)

		rtr.Get("/ping", resGP)
		rtr.Delete("/api/user/urls", resDAUU)

		MyHandler.ResCreateSQLTable = MyHandler.CreateSQLTable(db)
		log.Println(reflect.TypeOf(MyHandler.ResCreateSQLTable))
	} else {
		var db *sql.DB
		resDAUU := MyHandler.DeleteFuncAPIUserURLs(mapPost, mapGet, db, MyHandler.ResHandParam.DataBaseDSN)
		rtr.Delete("/api/user/urls", resDAUU)
	}

	if !MyHandler.ResHandParam.EnableHTTPS {
		log.Println("1")
		err := http.ListenAndServe(MyHandler.ResHandParam.PortNumber, rtr)
		if err != nil {
			log.Println(srError)
		}
	} else {
		log.Println("2")
		err := http.ListenAndServeTLS(MyHandler.ResHandParam.PortNumber, "../../internal/transport/server.cert", "../../internal/transport/server.key", rtr)
		if err != nil {
			log.Println(srError)
		}
	}

}
