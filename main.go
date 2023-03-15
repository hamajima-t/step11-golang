// STEP11: 集計ページの作成

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/tenntenn/sqlite"
)

func main() {

	// データベースへ接続
	// ドライバにはSQLiteを使って、
	// accountbook.dbというファイルでデータベース接続を行う
	db, err := sql.Open("sqlite", "accountbook.db")
	if err != nil {
		log.Fatal(err)
	}

	// AccountBookをNewAccountBookを使って作成
	ab := NewAccountBook(db)

	// テーブルを作成
	if err := ab.CreateTable(); err != nil {
		log.Fatal(err)
	}

	// HandlersをNewHandlersを使って作成
	hs := NewHandlers(ab)

	// ハンドラの登録
	http.HandleFunc("/", hs.ListHandler)
	http.HandleFunc("/ws", hs.WsEndpoint)
	http.HandleFunc("/save", hs.SaveHandler)
	http.HandleFunc("/summary", hs.SummaryHandler)

	// 静的ファイル（JavaScript/CSS）を読み込むために
	// ファイルサーバーを起動する
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	go hs.ListenToWsChannel()

	fmt.Println("http://localhost:8080 で起動中...")
	// HTTPサーバを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
