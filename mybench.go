package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Hello from my app")
	request := flag.Int64("n", 1, "number of requests")
	concurrency := flag.Int64("c", 1, "number of requests to perform at once")
	fmt.Println(request, concurrency)
	flag.Parse()
	flag.PrintDefaults()
}
