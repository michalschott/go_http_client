package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func MakeRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	secs := time.Since(start).Seconds()
	if err != nil {
		ch <- fmt.Sprintf("%.2f elapsed with error: %s", secs, err)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
	}
}

func main() {
	start := time.Now()
	ch := make(chan string)
	defer close(ch)
	sessions, _ := strconv.Atoi(os.Args[1])
	for c := 0; c < sessions; c++ {
		go MakeRequest(os.Args[2], ch)
	}

	for c := 0; c < sessions; c++ {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
