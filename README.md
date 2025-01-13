# NePP CTF

CTFの問題サイトはGitHub Pages（以下、Pages）で公開しています。

問題を解く人には本リポジトリの内容は見せず、Pagesの内容のみを共有するようにしてください（GitHub Organizationのプラン料金の都合でリポジトリを丸ごと公開しています）。

本READMEでは問題の作問者向けに利用技術等を紹介します。

## 利用技術
問題の作成にあたり利用している技術は以下です

- Pages（問題サイト）作成
  - [mdBook](https://rust-lang.github.io/mdBook/index.html)

### mdBook
Rust製の静的サイトジェネレーターです。`mdbook`ディレクトリ内に問題文を記述したマークダウンファイル等を配置しています。

`book.toml` は `mdBook` によって指定された設定ファイルです。一部の問題のフラグはmdBookの機能によってPages内に埋め込んでいますが、 `book.toml` から追うことができます。

[Renderers - mdBook Documentation](https://rust-lang.github.io/mdBook/format/configuration/renderers.html)

PagesへのデプロイはGitHub公式で用意されているActionsを利用しています。 `.github/workflows/mdbook.yml` を参照してください
