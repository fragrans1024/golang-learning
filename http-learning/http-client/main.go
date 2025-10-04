package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// http.Get使用var DefaultClient = &Client{}
	// DefaultClient未定义Transport，使用DefaultTransport
	// DefaultTransport显示开启TCP keepalive，并调整了参数
	// DefaultTransport未定义DisableKeepAlives，则为false，request/response后，保留TCP连接
	resp, err := http.Get("http://127.0.0.1:9001")
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
