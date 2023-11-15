# ビルドステージ
# 使用するGo言語のイメージを指定
FROM golang:1.21.4 AS builder

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードの依存関係をコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY *.go ./

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o twifixbot

# 実行ステージ
# 軽量なベースイメージを使用
FROM alpine:latest

# 作業ディレクトリを設定
WORKDIR /root/

# ビルドステージから実行可能ファイルをコピー
COPY --from=builder /app/twifixbot .

# 実行コマンド
CMD ["/root/twifixbot"]