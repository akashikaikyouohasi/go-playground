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
	debugLog("Received request for /users, method: %s", r.Method)

	switch r.Method {
	case http.MethodGet: // https://cs.opensource.google/go/go/+/refs/tags/go1.24.5:src/net/http/method.go;l=10
		// GETリクエストの処理
		debugLog("Handling GET request for /users")
		// レスポンスヘッダーを設定
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// ユーザー情報をJSON形式で返す
		jsonBlock, _ := json.Marshal(users)
		debugLog("JSON Response: %s", string(jsonBlock))

		// エラーハンドリング

		// レスポンスにJSONデータを書き込む
		fmt.Fprint(w, string(jsonBlock))
	case http.MethodPost:
		// POSTリクエストの処理
		debugLog("Handling POST request for /users")
		// レスポンスヘッダーを設定
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// リクエストボディを取得
		contentLength := r.ContentLength
		body := make([]byte, contentLength) // スライスを作成。
		r.Body.Read(body) // r.Body.Readは、リクエストボディからデータを読み取る
		debugLog("Request Body: %s", string(body))

		// ユーザー情報をデコード
		var newUser User
		if err := json.Unmarshal(body, &newUser); err != nil {
			// エラーハンドリング
			http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest) // 400 Bad Request。HTTPリクエストの処理中のエラー
			debugLog("Error unmarshalling JSON: %v", err) // %vは、デフォルトのフォーマットでの表現を出力する
			return
		}
		debugLog("New User: %+v", newUser) // %+vは構造体の場合にフィールド名を出力する
		users_string, _ := json.Marshal(users)
		debugLog("Users before addition: %v", string(users_string))

		// ユーザー追加
		newUser.Id = len(users) + 1
		users = append(users, newUser)
		
		users_string, _ = json.Marshal(users)
		debugLog("Users after addition: %v", string(users_string))


		// 201 Created を返す
		// https://cs.opensource.google/go/go/+/refs/tags/go1.24.5:src/net/http/status.go;l=9
		w.WriteHeader(http.StatusCreated)

		// レスポンスにJSONデータを書き込む
		jsonBlock, _ := json.Marshal(newUser)
		debugLog("JSON Response: %s", string(jsonBlock))
		fmt.Fprint(w, string(jsonBlock))
	}
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