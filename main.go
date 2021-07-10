package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/junichiseki0831/go-tech-blog/handler"
	"github.com/junichiseki0831/go-tech-blog/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sqlx.DB

// createMux関数の戻り値を格納
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)

	// URLと `handler.ArticleIndex` という処理を結びつける
	e.GET("/", handler.ArticleIndex)
	//ルーティング設定追加
	e.GET("/new", handler.ArticleNew)
	e.GET("/:id", handler.ArticleShow)
	e.GET("/:id/edit", handler.ArticleEdit)
	// Webサーバーをポート番号 8080 で起動する
	e.Logger.Fatal(e.Start(":8080"))
}

// アプリケーションインスタンスの生成
func createMux() *echo.Echo {

	// アプリケーションインスタンスを生成
	e := echo.New()

	// アプリケーションに各種ミドルウェアを設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	//cssの使用
	e.Static("/css", "src/css")
	//jsの使用
	e.Static("/js", "src/js")

	// アプリケーションインスタンスを返す
	return e
}

// DB接続
func connectDB() *sqlx.DB {
	//dsn := os.Getenv("DSN")
	//環境変数を設定していないためとりあえず直書き
	dsn := "sample_user:sample_user..@tcp(192.168.96.2:3306)/techblog"
	fmt.Println(dsn)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}
