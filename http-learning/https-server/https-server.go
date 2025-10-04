package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	https_enable_keylog()

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

// https://blog.csdn.net/orlobl/article/details/134871444
func https_enable_keylog() {
	keyLogFile, err := os.OpenFile("./keylog", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open keylog fail, err =", err)
		return
	}

	tlsConfig := tls.Config{
		KeyLogWriter: keyLogFile,
	}

	server := http.Server{
		Addr:      "127.0.0.1:443",
		Handler:   nil, // 使用DefaultServeMux
		TLSConfig: &tlsConfig,
	}

	http.HandleFunc("/", indexHandler)

	err = server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Println("ListenAndServeTLS fail, err =", err)
	}
}

// http.ListenAndServeTLS
// 将第一个参数addr和第4个参数handler封装为Server
// 然后调用server.ListenAndServeTLS(certFile, keyFile)
// func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error {
// 	server := &Server{Addr: addr, Handler: handler}
// 	return server.ListenAndServeTLS(certFile, keyFile)
// }

// net.(*Server) ListenAndServeTLS
// 将Server中的参数address通过net.Listen()转换为net.Listener，此处为TCP监听
// 然后调用net.(*Server) ServeTLS
// func (s *Server) ListenAndServeTLS(certFile, keyFile string) error {
// 	ln, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		return err
// 	}

// 	defer ln.Close()

// 	return s.ServeTLS(ln, certFile, keyFile)
// }

// net.(*Server) ServeTLS
// 处理TLSConfig
// 将TCP监听和TLSconfig组装成为TLS listener
// 最后调用net.(*Server) Serve(l net.Listener)
// func (s *Server) ServeTLS(l net.Listener, certFile, keyFile string) error {

// 	config := cloneTLSConfig(s.TLSConfig)
//  config.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)

// 	tlsListener := tls.NewListener(l, config)
// 	return s.Serve(tlsListener)
// }
