package model

// Article ...
type Article struct {
	// メタ情報を付与することでSQL実行と連携
	ID    int    `db:"id"`
	Title string `db:"title"`
}
