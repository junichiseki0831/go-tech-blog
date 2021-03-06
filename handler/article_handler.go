package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/junichiseki0831/go-tech-blog/model"
	"github.com/junichiseki0831/go-tech-blog/repository"
	"github.com/labstack/echo/v4"
)

// ArticleIndex ...
func ArticleIndex(c echo.Context) error {
	// 記事データの一覧を取得する
	articles, err := repository.ArticleList()
	if err != nil {
		log.Println(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// htmlに渡すデータ作成
	data := map[string]interface{}{
		//"Message":  "Article Index",
		"Message":  "Article Index Updated",
		"Now":      time.Now(),
		"Articles": articles, // 記事データをテンプレートエンジンに渡す
	}
	return render(c, "article/index.html", data)
}

// ArticleNew ...
func ArticleNew(c echo.Context) error {
	data := map[string]interface{}{
		"Message": "Article New",
		"Now":     time.Now(),
	}

	return render(c, "article/new.html", data)
}

// ArticleShow ...
func ArticleShow(c echo.Context) error {
	//　パスパラメータを抽出
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Show",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/show.html", data)
}

// ArticleEdit ...
func ArticleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Message": "Article Edit",
		"Now":     time.Now(),
		"ID":      id,
	}

	return render(c, "article/edit.html", data)
}

// ArticleCreateOutput ...
type ArticleCreateOutput struct {
  Article          *model.Article
  Message          string
  ValidationErrors []string
}

// ArticleCreate ...
func ArticleCreate(c echo.Context) error {
  // 送信されてくるフォームの内容を格納する構造体を宣言
  var article model.Article

  // レスポンスとして返却する構造体を宣言
  var out ArticleCreateOutput

  // フォームの内容を構造体に埋め込み
  if err := c.Bind(&article); err != nil {
    // エラーの内容をサーバーのログに出力
    c.Logger().Error(err.Error())

    // リクエストの解釈に失敗した場合は 400 エラーを返却
    return c.JSON(http.StatusBadRequest, out)
	}
	
	// バリデーションチェックを実行
	if err := c.Validate(&article); err != nil {
		// エラーの内容をサーバーのログに出力
		c.Logger().Error(err.Error())

		// エラー内容を検査してカスタムエラーメッセージを取得します。
		out.ValidationErrors = article.ValidationErrors(err)

		// 解釈できたパラメータが許可されていない値の場合は 422 エラーを返却
		return c.JSON(http.StatusUnprocessableEntity, out)
	}

  // repository を呼び出して保存処理を実行
  res, err := repository.ArticleCreate(&article)
  if err != nil {
    // エラーの内容をサーバーのログに出力
    c.Logger().Error(err.Error())

    // サーバー内の処理でエラーが発生した場合は 500 エラーを返却
    return c.JSON(http.StatusInternalServerError, out)
  }

  // SQL 実行結果から作成されたレコードの ID を取得
  id, _ := res.LastInsertId()

  // 構造体に ID をセット
  article.ID = int(id)

  // レスポンスの構造体に保存した記事のデータを格納
  out.Article = &article

  // 処理成功時はステータスコード 200 でレスポンスを返却
  return c.JSON(http.StatusOK, out)
}
