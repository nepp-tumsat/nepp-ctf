# [Level1] Dockerイメージのコンテナ内シェル実行

`ghcr.io/nepp-tumsat/nepp-ctf-linux-docker-q2`はNePP CTFによって公開されているDockerイメージです。

本イメージを実行したコンテナの `/nepp-ctf/flag.txt` ファイルからフラグを取得してください

## フラグ形式

`nepp{linux-docker-q2-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}`

## ヒント
イメージはdebianベースとなっており `bash` シェルが搭載されているためフラグの取得にはいくつかの方法があります

`docker container run` コマンドの公式ドキュメントを参考にしたください 

[docker container run — Docker-docs-ja 24.0 ドキュメント](https://docs.docker.jp/engine/reference/commandline/container_run.html)
