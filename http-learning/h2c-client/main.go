package main

import (
	"context"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLSContext: func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
				dialer := &net.Dialer{}
				return dialer.DialContext(ctx, network, addr)
			},
		},
	}
	resp, err := client.Get("http://127.0.0.1:8080")
	if err != nil {
		log.Println("http.Get fail, err =", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("read resp.Body fail, err =", err)
		return
	}

	log.Println(string(body))
}
