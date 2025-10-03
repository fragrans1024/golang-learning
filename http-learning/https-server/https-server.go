package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServeTLS("127.0.0.1:443", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("ListenAndServeTLS fail, err =", err)
	}
}

// 生成私钥
// openssl genrsa -out server.key 2048

// 用私钥生成证书申请文件csr。会进入交互模式（填写地区等信息）
// openssl req -new -key server.key -out server.csr

// 用私钥对证书申请进行签名从而生成证书
// openssl x509 -req -in server.csr -out server.crt -signkey server.key -days 3650
