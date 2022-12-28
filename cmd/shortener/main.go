package main

import (
	//"bufio"
	//"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	//"strconv"
	"math/rand"
	"time"
)

//var count int = 0

const portNumber = ":8080"

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Myfunc(mm_p map[string]string, mm_g map[string]string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			bp, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("Error 11")
				w.WriteHeader(400)
			} else {
				rnd_res := randSeq(6)
				mm_p[string(bp)] = rnd_res
				mm_g[rnd_res] = string(bp)

				result_post := "http://localhost:8080/" + rnd_res

				//fmt.Println(result_post)

				w.WriteHeader(201)
				w.Write([]byte(result_post))

				//count +=1
			}
		} else if r.Method == http.MethodGet {

			ug := r.URL.Path
			buf := strings.Replace(ug, "/", "", -1)
			out := string(buf)
			fmt.Println(out)
			buf1 := mm_g[out]
			fmt.Println("1")
			if mm_g[out] != "" {
				fmt.Println("2")
				w.Header().Add("Location", buf1)
				w.WriteHeader(http.StatusTemporaryRedirect)
				//w.Write([]byte(mm_g[out]))
				//fmt.Println(w.Header()["Location"])
				fmt.Println("3")
			} else {
				w.WriteHeader(400)
			}

		} else {
			fmt.Println("Error method")
			w.WriteHeader(400)
		}
	}
}

func main() {
	map_post := make(map[string]string)
	map_get := make(map[string]string)

	res := Myfunc(map_post, map_get)

	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", res)

	http.HandleFunc("/POST", res)

	//http.HandleFunc("/GET/", res)

	fmt.Printf("Starting application on port %v\n", portNumber)

	http.ListenAndServe(portNumber, nil)

}
