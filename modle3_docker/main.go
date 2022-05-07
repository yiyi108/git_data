package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Print(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	r.Header.Set("go", "hello")
	io.WriteString(w, r.Header.Get("go"))
	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	envs := os.Environ()
	for _, env := range envs {
		io.WriteString(w, env)
	}
	io.WriteString(w, os.Getenv("VERSION"))
	fmt.Println("version", os.Getenv("VERSION"))
	// 3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	log.Println(r.RemoteAddr + http.StatusText(200))
	// 4.当访问 localhost/healthz 时，应返回200
	//上面执行出来的返回是200
}

