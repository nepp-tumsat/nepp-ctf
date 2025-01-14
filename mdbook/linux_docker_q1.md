# [Level1] Dockerイメージのコンテナ実行

`ghcr.io/nepp-tumsat/nepp-ctf-linux-docker-q1`はNePP CTFによって公開されているDockerイメージです。

本イメージをコンテナ実行しフラグを取得してください

## フラグ形式

`nepp{linux-docker-q1-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}`

## ヒント
フラグはDockerfileの [ENTRYPOINT](https://docs.docker.jp/v1.12/engine/reference/builder.html#entrypoint) にて `echo` で出力されるため、イメージをコンテナ実行する際は引数指定不要です。
