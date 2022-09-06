# 2020/10/14最新versionを取得
FROM golang:1.19.0-alpine
# アップデートとgitのインストール！！
RUN apk update && apk add git curl
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
# ADD ./go/ /go/src/app
# ホットリロード追加
# RUN go install github.com/cosmtrek/air@v1.27.3
# CMD ["air", "-c", ".air.toml"]
