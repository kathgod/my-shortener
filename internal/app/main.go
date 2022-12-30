package handler

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetFunc(mm_p map[string]string, mm_g map[string]string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ug := r.URL.Path
		buf := strings.Replace(ug, "/", "", -1)
		out := string(buf)
		//fmt.Println(out)
		if mm_g[out] != "" {
			w.Header().Set("Location", mm_g[out])
			w.WriteHeader(307)
		} else {
			w.WriteHeader(400)
		}
	}
}

func PostFunc(mm_p map[string]string, mm_g map[string]string) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		bp, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error 11")
			w.WriteHeader(400)
		} else {
			rnd_res := RandSeq(6)
			mm_p[string(bp)] = rnd_res
			mm_g[rnd_res] = string(bp)
			result_post := "http://localhost:8080/" + rnd_res
			w.WriteHeader(201)
			w.Write([]byte(result_post))
		}
	}
}

func NAMfunc() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Error method")
		w.WriteHeader(400)
	}
}
