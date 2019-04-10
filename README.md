# INTERN_TEST
インターン選考の課題のリポジトリです。
`Golang`,`postgreSQL`,`docker-compose`を用いてRESTAPIを実装しました。

## 仕様
### エンドポイント

```
GET    /            # Hello Worldを表示します。
GET    /users       # user の一覧を表示
GET    /users/:id   # 指定した id の user を表示
POST   /users       # user を追加
PUT    /users/:id   # 指定した id の user を更新
DELETE /users/:id   # 指定した id の user を削除
```
#### `GET /`

"Hello world!!" を返します。

##### リクエストパラメータ
なし

##### レスポンスの例
HTTP ステータスコード 200

```
{"message":"Hello World!!"}⏎
```

#### `POST /users`

user を作成します。

##### リクエストパラメータ
|Name|Description|
|---|---|
|name|ユーザー名|
|email|Emailアドレス|

##### リクエストの例
```
$ curl -XPOST -H 'Content-Type:application/json' http://localhost:8080/users -d '{"name": "test", "email":"hoge@example.com" }'
```

##### レスポンス 
HTTP ステータスコード 200

```json
[
  {
    "id": 1,
    "name": "test",
    "email": "hoge@example.com",
    "created_at": "2019-01-23T21:25:27.246092+09:00",
    "updated_at": "2019-01-23T21:25:27.246092+09:00"
  }
]
```

#### `GET /users/:id`

id を指定して user を取得します。

##### リクエストパラメータ
なし

##### リクエストの例

```sh
curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users/1
```

##### レスポンスの例 
HTTP ステータスコード 200

```json
{
  "id": 1,
  "name": "onsd",
  "email": "onsd@example.com",
  "created_at": "0001-01-01T09:18:59+09:18",
  "updated_at": "2019-01-23T21:27:47.101667+09:00"
}
```


##### エラーレスポンスの例
指定した user が存在しなかった場合以下のレスポンスを返します。

HTTP ステータスコード 404

```json
{
  "error": "record not found"
}
```

#### `GET /users`

user の一覧を取得します。

リクエスト

```sh
$ curl -H 'Content-Type:application/json' http://localhost:8080/users
```

#### `PUT /users/:id`

id で指定した user の名前、メールアドレスを更新します。

##### リクエストパラメータ
|Name|Description|
|---|---|
|name|ユーザー名|
|email|Emailアドレス|


```sh
$ curl -XPUT -H 'Content-Type:application/json' http://localhost:8080/users/1 -d '{"name": "onsd", "email": "onsd@example.com" }'
```

##### レスポンスの例
HTTP ステータスコード 200

```json
{
  "id": 1,
  "name": "onsd",
  "email": "onsd@example.com",
  "created_at": "2019-01-23T19:10:32.890846+09:00",
  "updated_at": "2019-01-23T12:31:06.1564117Z"
}
```

##### エラーレスポンスの例
指定した id のユーザが存在しなかった場合以下のレスポンスを返します。

HTTP ステータスコード 404


```json
{
  "error": "record not found"
}
```

#### `DELETE /users/:id`

id で指定した user を削除します。

##### リクエストの例

```sh
$ curl -XDELETE -H 'Content-Type:application/json' http://localhost:8080/users/1
```

##### レスポンスの例
HTTP ステータスコード 204

レスポンスはありません。

```json

```

##### エラーレスポンスの例
指定した id のユーザが存在しなかった場合以下のレスポンスを返します。

HTTP ステータスコード 409


```json
{
  "error": "record not found"
}
```