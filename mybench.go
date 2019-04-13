package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type responseInfo struct {
	status   int
	bytes    int64
	duration time.Duration
}

type summaryInfo struct {
	requested int64
	responded int64
}

func checkLink(link string, c chan responseInfo) {
	start := time.Now()
	res, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	read, _ := io.Copy(ioutil.Discard, res.Body)
	c <- responseInfo{
		status:   res.StatusCode,
		bytes:    read,
		duration: time.Now().Sub(start),
	}
}

func main() {
	fmt.Println("Hello from my app")
	request := flag.Int64("n", 1, "number of requests")
	concurrency := flag.Int64("c", 1, "number of requests to perform at once")
	fmt.Println(request, concurrency)
	flag.Parse()

	if flag.NArg() == 0 || *request == 0 || *request < *concurrency {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	link := flag.Arg(0)
	c := make(chan responseInfo)
	summary := summaryInfo{}

	for i := int64(0); i < *concurrency; i++ {
		summary.requested++
		go checkLink(link, c)
	}

	for response := range c {
		if summary.requested < *request {
			summary.requested++
			go checkLink(link, c)
		}
		summary.responded++
		fmt.Println(response)
		if summary.responded == summary.requested {
			break
		}
	}

}
