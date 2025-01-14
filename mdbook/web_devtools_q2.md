# [Level1] HTMLページのネットワークを読み取る

本HTMLページ <https://nepp-tumsat.github.io/nepp-ctf/> ではフラグを持つ外部リソースをJavaScriptコード実行によるネットワーク経由で取得しています。

ブラウザのデベロッパーツール等を活用しフラグを探してください

## フラグ形式

`nepp{web-devtools-q2-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}`

## ヒント

現代のWebアプリケーションの主流となっている[シングルページアプリケーション（SPA）](https://developer.mozilla.org/ja/docs/Glossary/SPA)では、外部リソースへのHTTPリクエストによるコンテンツデータの取得が頻繁に行われます。

JavaScript経由で行われるHTTPリクエストのリクエスト・レスポンス情報は、デベロッパーツールのネットワークタブから検索が可能です