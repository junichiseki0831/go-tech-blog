package model

import "time"

// Article ...
type Article struct {
	// メタ情報を付与することでSQL実行と連携
	ID      int       `db:"id" form:"id"`
	Title   string    `db:"title" form:"title"`
	Body    string    `db:"body" form:"body"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}
