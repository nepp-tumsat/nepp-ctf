# NePP CTF

CTFの問題サイトはGitHub Pages（以下、Pages）で公開しています。

問題を解く人には本リポジトリの内容は見せず、Pagesの内容のみを共有するようにしてください（GitHub Organizationのプラン料金の都合でリポジトリを丸ごと公開しています）。

本READMEでは問題の作問者向けに利用技術等を紹介します。

# ディレクトリ構成・役割

```
❯ tree -L 2
.
├── README.md
├── .github
│   └── workflows
├── book
│   ├── ...
├── book.toml
├── mdbook
│   ├── SUMMARY.md
│   ├── introduction.md
│   ├── linux_docker.md
│   ├── linux_docker_q1.md
│   ├── web_devtools.md
│   ├── web_devtools_q1.md
│   └── ...
├── questions
|   ├── linux_docker
|   └── web_devtools
└── server
    └── k8s
```

- `mdbook`, `bool.toml`, `book`
  - [問題サイトの作成](https://github.com/nepp-tumsat/nepp-ctf/wiki/%E5%95%8F%E9%A1%8C%E3%82%B5%E3%82%A4%E3%83%88%E4%BD%9C%E6%88%90)
- `questions`
  - [作問実装・解答のディレクトリ配置](https://github.com/nepp-tumsat/nepp-ctf/wiki/%E4%BD%9C%E5%95%8F%E5%AE%9F%E8%A3%85%E3%81%AE%E3%83%87%E3%82%A3%E3%83%AC%E3%82%AF%E3%83%88%E3%83%AA%E9%85%8D%E7%BD%AE)
-  `.github/workflows`
  -  [問題サイトのデプロイ方法](https://github.com/nepp-tumsat/nepp-ctf/wiki/%E5%95%8F%E9%A1%8C%E3%82%B5%E3%82%A4%E3%83%88%E3%81%AE%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E6%96%B9%E6%B3%95)
  -  [作問Dockerイメージのデプロイ方法](https://github.com/nepp-tumsat/nepp-ctf/wiki/%E4%BD%9C%E5%95%8FDocker%E3%82%A4%E3%83%A1%E3%83%BC%E3%82%B8%E3%81%AE%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E6%96%B9%E6%B3%95)
- `server/k8s`
  - [問題サーバーの実装](https://github.com/nepp-tumsat/nepp-ctf/wiki/%E5%95%8F%E9%A1%8C%E3%82%B5%E3%83%BC%E3%83%90%E3%83%BC%E3%81%AE%E5%AE%9F%E8%A3%85)   
