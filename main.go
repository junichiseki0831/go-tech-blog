package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// テンプレートファイルを配置するディレクトリへの相対パス格納
const tmplPath = "src/template/"

// createMux関数の戻り値を格納
var e = createMux()

func main() {

	// URLと `articleIndex` という処理を結びつける
	e.GET("/", articleIndex)
	//ルーティング設定追加
	e.GET("/new", articleNew)
	e.GET("/:id", articleShow)
	e.GET("/:id/edit", articleEdit)
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

// 一覧ページの生成
func articleIndex(c echo.Context) error {
	// ステータスコード 200 で、"Hello, World!" という文字列をレスポンス
	//return c.String(http.StatusOK, "Hello, World!")

	// htmlに渡すデータ作成
	data := map[string]interface{}{
		"Message": "Article Index",
		"Now":     time.Now(),
	}

	return render(c, "article/index.html", data)
}

func articleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return render(c, "article/new.html", data)
}

func articleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/show.html", data)
}

func articleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/edit.html", data)
}

// pongo2でテンプレートファイルとデータから HTML を生成
func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
}

func render(c echo.Context, file string, data map[string]interface{}) error {

	// 定義した htmlBlob() 関数を呼び出し、生成された HTML をバイトデータとして受け取る
	b, err := htmlBlob(file, data)
	// エラーチェック
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	// ステータスコード 200 で HTML データをレスポンス
	return c.HTMLBlob(http.StatusOK, b)
}
