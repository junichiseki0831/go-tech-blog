FROM golang:1.16.3-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# appディレクトリの作成
RUN mkdir /go-tech-blog
# ワーキングディレクトリの設定
WORKDIR /go-tech-blog
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go-tech-blog

EXPOSE 8080