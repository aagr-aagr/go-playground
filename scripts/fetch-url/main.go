package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("couldn't fetch the url: %v", err)
	}
	defer resp.Body.Close()
	fmt.Printf("The status code is: %v \n", resp.StatusCode)
	body, err := io.Copy(io.Discard, resp.Body)

	if err != nil {
		log.Fatalf("Couldn't read the response: %v \n", err)
	}

	fmt.Printf("The response size is: %v \n", body)

}
