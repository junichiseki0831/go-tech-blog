FROM golang:1.16.3-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go-tech-blog
# ホットリロード追加
RUN go get github.com/pilu/fresh
RUN go get -u -v bitbucket.org/liamstask/goose/cmd/goose
# ワーキングディレクトリの設定
WORKDIR /go-tech-blog
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go-tech-blog