package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler ( w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Go Playground!!")
}

func handler_hello_world ( w http.ResponseWriter, r *http.Request) {
	// レスポンスヘッダーを設定
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// 文字列を返す
	fmt.Fprint(w, "Hello, World!!!!")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello_world", handler_hello_world)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())

}