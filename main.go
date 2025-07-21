package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
)

var debugMode = os.Getenv("DEBUG") == "true"

// Debugログ出力用の関数
func debugLog(format string, value ...interface{}) { //interface{}は任意の型
	if debugMode {
		log.Printf("[DEBUG] " + format, value... )
	}
}

// ユーザー定義
type User struct {
	Id int `json:"id"` //構造体タグあり
	Name string `json:"name"`
}

var users = []User{
	{ Id: 1, Name: "Alice" },
	{ Id: 2, Name: "Bob" },
}

func handler ( w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome Go Playground!!")
}

func handler_hello_world ( w http.ResponseWriter, r *http.Request) {
	// レスポンスヘッダーを設定
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// 文字列を返す
	fmt.Fprint(w, "Hello, World!!!!")
}

func handler_users ( w http.ResponseWriter, r *http.Request) {
	// レスポンスヘッダーを設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// ユーザー情報をJSON形式で返す
	jsonBlock, _ := json.Marshal(users)
	debugLog("JSON Response: %s", string(jsonBlock))

	// エラーハンドリング

	// レスポンスにJSONデータを書き込む
	fmt.Fprint(w, string(jsonBlock))
}

func main() {
	var httpServer http.Server
	
	// ルーティング設定
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello_world", handler_hello_world)
	http.HandleFunc("/users", handler_users)


	// サーバーの設定
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}