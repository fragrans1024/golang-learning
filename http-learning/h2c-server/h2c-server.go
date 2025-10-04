package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// https://pkg.go.dev/golang.org/x/net/http2/h2c

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})
	h2s := &http2.Server{
		// ...
	}
	h1s := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(handler, h2s),
	}
	log.Fatal(h1s.ListenAndServe())
}

// 使用curl测试
// curl --http2-prior-knowledge http://127.0.0.1:8080
// curl --http2                 http://127.0.0.1:8080
// curl                         http://127.0.0.1:8080
