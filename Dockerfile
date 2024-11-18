FROM golang:1.23.3 AS build

WORKDIR /workdir

COPY . .

# 依存関係のインストール（開発環境）
# install opusfile
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -y \
    pkg-config libopusfile-dev=0.12-4 \
    && rm -rf /var/lib/apt/lists/*

# `app/`ディレクトリが存在する場合、`-o app`を指定すると`app/${project_name}`という実行ファイルが生成される
# ディレクトリにならなさそうなapp.exeにしておく
RUN go build -ldflags "-s -w" -trimpath -mod=vendor -o app.exe .

FROM ubuntu:22.04

WORKDIR /workdir

# set timezone to Asia/Tokyo
# install ca-certificates
# install opusfile
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive \
    apt-get install -y \
    tzdata \
    ca-certificates \
    pkg-config libopusfile0=0.9+20170913-1.1build1 \
    && rm -rf /var/lib/apt/lists/* \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

COPY --from=build /workdir/app.exe app.exe
CMD ["./app.exe"]
