package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Get("https://127.0.0.1")
	if err != nil {
		fmt.Println("http.GET fail, err =", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body fail, err =", err)
		return
	}

	fmt.Println("body =", string(body))

	time.Sleep(120 * time.Second)
}
