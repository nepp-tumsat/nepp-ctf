# Network - HTTP

OSI参照モデルのL7　アプリケーション層として一般的に使われているHTTPプロトコルについて学びます。

NePP CTFで用意されている問題サーバー（ <https://ctf.nepp-tumsat.dev/> ）にリクエストを送りHTTPの基本を身につけます

## 推奨ツール
HTTPリクエストはlinuxの `curl` コマンドで行うのが一般的ですが、より便利なツールとして[HTTPie](https://httpie.io/)等があります。

## 準備
問題サーバーは以下のGitHub Actionsからの起動が必要です。最新のコミット分のワークフローをRerunすることでGKEへのApplyが完了します

<https://github.com/nepp-tumsat/nepp-ctf/actions/workflows/google.yml>