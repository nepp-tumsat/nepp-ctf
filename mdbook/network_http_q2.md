# [Level1] POSTメソッドによるJSONデータ送信

<https://ctf.nepp-tumsat.dev/network-http-q2> エンドポイントへのPOSTリクエストを行うことでフラグを取得できます

以下のJSON形式のデータを送信してください、i, j, kに割り当てるのは数値型の値です

```json
{
    "meiji_maru": 0, // number type
    "marine_cafe": 0, // number type
    "kaio_dormitory": 0 // number type
}
```

JSONフィールドの数値は [アクセス | 東京海洋大学について | 国立大学法人 東京海洋大学](https://www.kaiyodai.ac.jp/overview/access/) ページにある越中島キャンバスの地図から、「明治丸」、「マリン・カフェ（ワールドマリン カフェ）」、「学生寮（海王寮）」に割り当てられている番号をそれぞれ対応する英名のフィールドに設定してください。

## フラグ形式

`nepp{network-http-q2-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx}`

## ヒント
JSONデータを送信する場合、`Content-Type` HTTPヘッダーを指定する必要があります
