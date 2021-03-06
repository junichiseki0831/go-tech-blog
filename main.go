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
	"gopkg.in/go-playground/validator.v9" //バリデーション
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
	e.POST("/", handler.ArticleCreate)
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
	// CSRF対策
	e.Use(middleware.CSRF())

	//cssの使用
	e.Static("/css", "src/css")
	//jsの使用
	e.Static("/js", "src/js")
	//バリデーション
	e.Validator = &CustomValidator{validator: validator.New()}

	// アプリケーションインスタンスを返す
	return e
}

// DB接続
func connectDB() *sqlx.DB {
	//dsn := os.Getenv("DSN")
	//環境変数を設定していないためとりあえず直書き
	dsn := "sample_user:sample_user..@tcp(192.168.112.2:3306)/techblog?parseTime=true"
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

  // CustomValidator ...
  type CustomValidator struct {
    validator *validator.Validate
  }
 
  // Validate ...
  func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
  }
