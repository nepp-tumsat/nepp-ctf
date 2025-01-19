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
└── questions
    ├── linux_docker
    └── web_devtools
```

## `mdbook`, `bool.toml`, `book`
[問題サイトの作成](https://github.com/nepp-tumsat/nepp-ctf/wiki/%E5%95%8F%E9%A1%8C%E3%82%B5%E3%82%A4%E3%83%88%E4%BD%9C%E6%88%90) を参照

## `questions` 作問
各作問実装が配置されたディレクトリです。

`questions` 直下に `ジャンル_題材` の形式のディレクトリを作成します。さらにその子ディレクトリ下に `q1`, `q2` のような形で作問別の実装を配置してください

また各問題ディレクトリには実装とは別に `answer.md` ファイルを残し、解答・解説・作問補足等を残しておくことを推奨します

## `.github/workflows` 問題サイト・問題パッケージのデプロイ
各ワークフローの概要について補足していきます

### `mdbook.yml`
mdBookビルトによって問題サイトをPagesにデプロイするワークフローです

`main`ブランチへのpush時にmdBook関連ファイルに変更があったケースがトリガーとなります。

### `docker-publish-linux-docker-{問題番号}.yml`
Linux Dockerジャンルにおいて作問実装のDockerfileを元にDockerイメージをGitHub Packagesへpushします

GitHub Organizationの無料プランではPackagesのストレージに上限があるため、無闇にトリガーされないよう `git tag` でトリガーを手動化しています。

Dockerイメージ作成時は該当gitコミットに対し、　

```
git tag v0.1.0-linux-docker-q1
git push origin v0.1.0-linux-docker-q1
```

のように `semver + 問題名` のタグ付け、タグのpushをしワークフローが起動するよう手動対応をしてください

