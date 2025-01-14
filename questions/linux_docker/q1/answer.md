# 解答

nepp{linux-docker-q1-bfbe2788-a7a3-4edb-b575-8c7d80f6782d}

# 解説
Dockerイメージの取得とコンテナ実行を行います

```
❯ docker image pull ghcr.io/nepp-tumsat/nepp-ctf-linux-docker-q1
Using default tag: latest
latest: Pulling from nepp-tumsat/nepp-ctf-linux-docker-q1
e474a4a4cbbf: Pull complete 
Digest: sha256:e2f2bef5035b32b86e2a4b667b773ca9b98a644225c98ea66ea64184488502e9
Status: Downloaded newer image for ghcr.io/nepp-tumsat/nepp-ctf-linux-docker-q1:latest
ghcr.io/nepp-tumsat/nepp-ctf-linux-docker-q1:latest

❯ docker container run ghcr.io/nepp-tumsat/nepp-ctf-linux-docker-q1
flag is nepp{linux-docker-q1-bfbe2788-a7a3-4edb-b575-8c7d80f6782d}
```

# 作問補足

Dockerコンテナ実行時のコマンド実行ではDockerfile内で `CMD` ではなく `ENTRYPOINT` を用います。

`ENTRYPOINT` と `CMD` の使い分けは以下が参考になります

[ENTRYPOINTは「必ず実行」、CMDは「（デフォルトの）引数」 ‣ Pocketstudio.Net](https://pocketstudio.net/2020/01/31/cmd-and-entrypoint/)

