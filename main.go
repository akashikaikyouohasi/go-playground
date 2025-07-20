package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler ( w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		// エラーがある場合の処理
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return 
	}
	fmt.Println(string(dump))
	fmt.Fprint(w, "<html><body>hello</body></html>")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())

}