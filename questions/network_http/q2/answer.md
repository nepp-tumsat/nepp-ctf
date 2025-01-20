# 解答

nepp{network-http-q2-f4972292-017c-4de5-be5d-b369750e0542}

# 解説
curl コマンドによって問題文にあるURLへHTTPリクエストを送ります

```
❯ curl -X POST https://ctf.nepp-tumsat.dev/network-http-q2 -H 'Content-Type: application/json' -d '{"meiji_maru":5,"marine_cafe":15,"kaio_dormitory":24}'
{"flag":"nepp{network-http-q2-f4972292-017c-4de5-be5d-b369750e0542}"}
```

HTTPieを利用する場合、以下のようなコマンドになります
```
http POST https://ctf.nepp-tumsat.dev/network-http-q2 meiji_maru=5 marine_cafe=5 kaio_dormitor=24
```
