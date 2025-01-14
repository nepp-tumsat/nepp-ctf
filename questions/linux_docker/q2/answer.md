# 解答

nepp{linux-docker-q2-4cb51b88-e1e0-4ad4-b59a-bf8c64591a09}

# 解説
ヒントにも書いた通り `docker container run` による実行方法によっていくつかの手段があります


1. `docker container run` の引数に `cat` を指定しコマンド一回で出力する方法
```
docker container run ghcr.io/nepp-tumsat/nepp-ctf-linux-docker-q2:latest cat /nepp-ctf/flag.txt
```

2. `docker container run` の引数に `bash` を指定しコンテナシェル上でフラグを探す方法
```
❯ docker container run -it --rm  ghcr.io/nepp-tumsat/nepp-ctf-linux-docker-q2:latest bash
root@5aaa0864adf3:/nepp-ctf# cat flag.txt
nepp{linux-docker-q2-4cb51b88-e1e0-4ad4-b59a-bf8c64591a09}root@5aaa0864adf3
```

`-it`によってコンテナ上でインタラクティブにコマンドが実行できます、 `--rm` の指定によってコンテナを抜けた後に自動でコンテナが削除されます

# 作問補足
本問題のDockerfileは `USER` を指定していないためroot権限でのコマンド実行が可能です。本来望ましくありませんが問題難易度を下げるためあえて権限を緩めています。

[Dockerfile のベストプラクティス — Docker-docs-ja 1.9.0b ドキュメント](https://docs.docker.jp/engine/articles/dockerfile_best-practice.html#user)