package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/krylphi/helloworld-request-sender/handler"
)

const (
	requests int = 1000000000
	threads  int = 1000
	//requests int = 1000
	//threads int = 100
)

func main() {
	env := func(key, def string) string {
		value, ok := os.LookupEnv(key)
		if !ok {
			return def
		}
		return value
	}
	endpoint := fmt.Sprintf("http://%s:%s/log", env("ADDR", "0.0.0.0"), env("PORT", "8902"))

	h := handler.NewHandler(endpoint, threads, requests)
	wg := sync.WaitGroup{}
	wg.Add(1)
	h.Handle()

}
