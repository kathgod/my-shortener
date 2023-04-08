package handler_test

import (
	"fmt"
	"log"
	"net/http"
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

	req2, err := http.NewRequest("GET", "shortURLnotExist", nil)
	if err != nil {
		log.Println(err)
	}

	status2, _ := h.LogicGetFunc(req2, handMapGet)
	fmt.Println(status2)
}

// Output:
//[307]
//[400]
