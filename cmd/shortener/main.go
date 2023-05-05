package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	MyHandler "urlshortener/internal/app"
	MyStorage "urlshortener/internal/database"
	MyLogic "urlshortener/internal/logic"
)

const (
	srError = "Server Error"
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

	MyLogic.ResHandParam.PortNumber = MyLogic.HandParam("SERVER_ADDRESS", srvAddress)
	MyLogic.ResHandParam.BaseURL = MyLogic.HandParam("BASE_URL", bsURL)
	MyLogic.ResHandParam.FileStoragePath = MyLogic.HandParam("FILE_STORAGE_PATH", flStoragePth)
	MyLogic.ResHandParam.DataBaseDSN = MyLogic.HandParam("DATABASE_DSN", datadbaseDsn)

	enableHTTPSBuff := MyLogic.HandParam("ENABLE_HTTPS", enableHTTPS)

	buff, err := strconv.ParseBool(enableHTTPSBuff)
	MyLogic.ResHandParam.EnableHTTPS = buff
	if err != nil {
		MyLogic.ResHandParam.EnableHTTPS = false
	}

	config := MyLogic.HandParam("CONFIG", configFile)
	if config != "" {
		MyLogic.HandConfigParam(config)
	}
	log.Println(MyLogic.ResHandParam)

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

	if MyLogic.ResHandParam.DataBaseDSN != "" {
		db := MyStorage.OpenDB(MyLogic.ResHandParam.DataBaseDSN)
		log.Println("Point 1")
		resGP := MyHandler.GetFuncPing(db)
		resDAUU := MyHandler.DeleteFuncAPIUserURLs(mapPost, mapGet, db, MyLogic.ResHandParam.DataBaseDSN)
		log.Println("Point 2")
		rtr.Get("/ping", resGP)
		rtr.Delete("/api/user/urls", resDAUU)
		log.Println("Point 3")
		MyLogic.ResCreateSQLTable = MyLogic.CreateSQLTable(db)
		log.Println(reflect.TypeOf(MyLogic.ResCreateSQLTable))
	} else {
		var db *sql.DB
		resDAUU := MyHandler.DeleteFuncAPIUserURLs(mapPost, mapGet, db, MyLogic.ResHandParam.DataBaseDSN)
		rtr.Delete("/api/user/urls", resDAUU)
	}

	server := &http.Server{Addr: MyLogic.ResHandParam.PortNumber, Handler: rtr}
	if !MyLogic.ResHandParam.EnableHTTPS {
		err := server.ListenAndServe()
		if err != nil {
			log.Println(srError)
		}
	} else {
		err := server.ListenAndServeTLS("../../internal/transport/server.cert", "../../internal/transport/server.key")
		if err != nil {
			log.Println(srError)
		}
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-sigChan
	log.Println("Received signal")
	shutdownctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = server.Shutdown(shutdownctx); err != nil {
		log.Fatalf("Error shutdown server: %v", err)
	}
}
