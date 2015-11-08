package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var urls  = []string {
	"http://www.google.com",
	"http://www.yahoo.com",
	"http://github.com",
	"http://bitbucket.org",
}

func main() {
	urlChannel := make(chan string)

	for _, url := range urls {
		go printIt(url, urlChannel)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s", <-urlChannel)
	}
}

func printIt(url string, urlChannel chan string) {
	length, err := getSize(url)
	if err != nil {
		os.Exit(1)
	}

	urlChannel <- fmt.Sprintf("Site %s has length %d \n", url, length)
}


func getSize(url string) (int, error) {
	response, err := http.Get(url)

	if err != nil {
		os.Exit(1)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil
}
