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
	status      int
	bytes       int64
	duration    time.Duration
	timeRequest time.Duration
}

type summaryInfo struct {
	requested int64
	responded int64
}

func checkLink(link string, c chan responseInfo) {
	start := time.Now()
	res, err := http.Get(link)
	timeRequest := time.Now().Sub(start)
	if err != nil {
		panic(err)
	}
	read, _ := io.Copy(ioutil.Discard, res.Body)
	c <- responseInfo{
		status:      res.StatusCode,
		bytes:       read,
		duration:    time.Now().Sub(start),
		timeRequest: timeRequest,
	}
}

func main() {

	request := flag.Int64("n", 1, "number of requests")
	concurrency := flag.Int64("c", 1, "number of requests to perform at once")
	timeOut := flag.Int64("timeout", 30, "the time out wait for each response")
	timeLimt := flag.Int64("timelimit", 30, "the time limit spend for benchmarking")

	flag.Parse()

	if flag.NArg() == 0 || *request == 0 || *request < *concurrency {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	timerOut := time.NewTimer(time.Duration(*timeOut) * time.Second)
	timerLimit := time.NewTimer(time.Duration(*timeLimt) * time.Second)
	link := flag.Arg(0)
	c := make(chan responseInfo)
	summary := summaryInfo{}
	begin := time.Now()
	for i := int64(0); i < *concurrency; i++ {
		summary.requested++
		go checkLink(link, c)

	}

	for response := range c {
		if summary.requested < *request {
			summary.requested++
			go checkLink(link, c)
		}
		select {
		case <-timerOut.C:
			fmt.Println("Responsed time exceed expectations")

		default:
			summary.responded++

			fmt.Println(response)
		}

		if summary.responded == summary.requested {
			break
		}
	}

	end := time.Now().Sub(begin)

	fmt.Println("Time taken for tests")
	fmt.Println(end)

	<-timerLimit.C

}
