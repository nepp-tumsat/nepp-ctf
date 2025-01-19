# [Level1] GETメソッドによるヘルスチェック

<https://ctf.nepp-tumsat.dev/health> エンドポイントへのGETリクエストによってフラグを取得してください

## フラグ形式

`nepp{network-http-q1-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}`

## 補足

ヘルスチェックは、Webアプリケーションのデプロイ時にアプリケーションが起動完了しているかを判別するために用いられるパターンです。アプリケーションサーバーを構築するときは最初に作成されるAPIエンドポイントになっています。