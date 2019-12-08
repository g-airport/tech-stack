// all goroutine s

// https://github.com/golang/go/issues/34457



package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func sayHello(w http.ResponseWriter, r *http.Request) {}

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Second * 10)
		}()
	}

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}