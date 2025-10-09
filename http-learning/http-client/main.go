package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	resp, err := http_client_with_trace()
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

func http_default_client() (resp *http.Response, err error) {
	// http.Get使用var DefaultClient = &Client{}
	// DefaultClient未定义Transport，使用DefaultTransport
	// DefaultTransport显示开启TCP keepalive，并调整了参数
	// DefaultTransport未定义DisableKeepAlives，则为false，request/response后，保留TCP连接
	resp, err = http.Get("http://127.0.0.1:9001")

	return resp, err
}

// go build -gcflags="all=-N -l"
// 显示设置client
// 显示设置Transport
// 如果Transport不设置DialContext/Dial，
// 根据net/http.(*Transport).dial()定义，使用net/http.zeroDialer.DialContext
// 虽然KeepAlive和KeepAliveConfig未配置
// 根据net.newTCPConn()和net.(*TCPConn).SetKeepAliveConfig()，开启keepalive，使用默认参数
// 如果配置KeepAlive小于0，则关闭tcp keepalive

// request/response后，不保留TCP连接的方法
// https://blog.csdn.net/cyberspecter/article/details/83308348

// https://pkg.go.dev/net/http/httptrace#example-package
func http_client_with_trace() (resp *http.Response, err error) {
	req, _ := http.NewRequest("GET", "http://127.0.0.1:9001", nil)

	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS info: %+v\n", dnsInfo)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	resp, err = http.DefaultTransport.RoundTrip(req)

	return resp, err
}
