package handler_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	h "urlshortener/internal/app"
)

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
