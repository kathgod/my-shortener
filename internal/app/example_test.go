package handler_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	h "urlshortener/internal/app"
)

func ExampleGetFunc() {
	handMapGet := map[string]string{
		"shortURL": "originalURL",
	}
	req1, err := http.NewRequest("GET", "shortURL", nil)
	if err != nil {
		log.Println(err)
	}

	status1, _ := h.LogicGetFunc(req1, handMapGet)
	fmt.Println(status1)

	req2, err := http.NewRequest("GET", "shortURLNotExist", nil)
	if err != nil {
		log.Println(err)
	}

	_, _ = h.LogicGetFunc(req2, handMapGet)
}

// Output:
//[307]
//[400]

func ExamplePostFunc() {
	handMapPost := map[string]string{}
	handMapGet := map[string]string{}

	req, err := http.NewRequest("POST", "/", strings.NewReader("originalURL"))
	if err != nil {
		log.Println(err)
	}
	nr := httptest.NewRecorder()

	status, _ := h.LogicPostFunc(nr, req, handMapPost, handMapGet)

	fmt.Println(status)

}

// Output:
//[201]
