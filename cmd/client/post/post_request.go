package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	//"net/url"
	"os"
	"strings"
)

func main() {
	//data := url.Values{}

	fmt.Println("Введите URL")
	reader := bufio.NewReader(os.Stdin)

	long, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	long = strings.TrimSuffix(long, "\n")

	//fmt.Println(long)

	//data.Set("url",long)

	client := &http.Client{}

	endpoint := "http://localhost:8080/POST"

	//fmt.Println(data.Encode())

	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBufferString(long))
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(2)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("ERROR 3")
		os.Exit(3)
	}
	fmt.Println(response.Status)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ERROR 4")
		os.Exit(4)
	}

	fmt.Println(string(body))
}
