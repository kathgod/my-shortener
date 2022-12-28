package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"io"
	"net/http"
	//"net/url"
	"os"
	"strings"
)

func main() {
	//data := url.Values{}

	fmt.Println("Введите ID")
	reader := bufio.NewReader(os.Stdin)

	ID, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}
	ID = strings.TrimSuffix(ID, "\n")

	//fmt.Println(ID)

	//data.Set("url",long)

	client := &http.Client{}

	endpoint := ID

	//fmt.Println(endpoint)

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		fmt.Println("ERROR 2")
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
