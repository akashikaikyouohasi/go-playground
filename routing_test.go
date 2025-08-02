package main
import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"reflect"
)

func TestHelloWorld(t *testing.T){
	req, err := http.NewRequest("GET", "/hello_world" ,nil) //http.NewRequestは、doしないとリクエストが送信されない
	fmt.Println(reflect.TypeOf(req)) // reqの型を表示する
	if err != nil {
		// errがnulでない場合は、エラーが発生したことになる
		t.Fatal(err) //Fatalは、処理が止まる
	}

	rr := httptest.NewRecorder() // http.ResponseWriterを満たす、*httptest.ResponseRecorderオブジェクトを取得する
	fmt.Println(reflect.TypeOf(rr)) // rrの型を表示する
	handler_hello_world(rr, req) // ハンドラー関数を呼び出して実行
	fmt.Printf("Response: %+v", rr) // レスポンスの内容を表示する

	if status := rr.Code; status != http.StatusOK {
		// ステータスコードが200ではない場合
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello, World!!!!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}