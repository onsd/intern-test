# INTERN_TEST
インターン選考の課題のリポジトリです。

`Golang`,`postgreSQL`,`docker-compose`を用いてRESTAPIを実装しました。

## 必要条件
+ [Docker](https://www.docker.com/) (Docker-compose が利用できること)
+ [Go](https://golang.org/) 1.12.1+ (ローカルで開発するために、Go modulesを使用できるようにしてください。)
## インストール

```sh
git clone https://github.com/onsd/intern-test
cd intern-test
docker-compose up
```

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

## DB接続情報について
docker-compose.yml,go/DockerfileにデフォルトのDB接続情報が書いてあります。


任意のDBを使用するためには、以下の環境変数を変更してください。

`docker-compose.yml`
```yaml
    environment:
      - HOSTNAME=postgres
      - USER=postgres
      - DBNAME=wantedly
      - PASSWORD=password
      - CGO_ENABLED=0
      - GO111MODULE=on
      - PORT=8080
      - DB_PORT=5432
 ```

 `go/Dockerfile`
 ```dockerfile
 ENV HOSTNAME postgres
ENV USER postgres
ENV DBNAME wantedly
ENV PASSWORD password
ENV CGO_ENABLED 0 
ENV GO111MODULE on
ENV PORT 8080
ENV DB_PORT 5432
```


## ローカルで開発するために
###  PostgreSQLをローカルで動作させる
`docker-compose.yml`を以下のように書き換えてください。
```
version: "3"
services:
  postgres:
    image: postgres
    container_name: postgres
    ports:
      - 5432:5432
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=password
      - TZ=`ls -la /etc/localtime | cut -d/ -f8-9`
    tty: true
    restart: always
    user: postgres
    volumes:
      - ./postgresql/init:/docker-entrypoint-initdb.d
      - postgres-db:/var/lib/postgresql/data
volumes:
  postgres-db:
    driver: local
```
もしくは、`develop/local`ブランチをチェックアウトすることで上記の`docker-compose.yml`に切り替えることができます。

そのあと、`docker-compose up`してください。



###  環境変数を設定する

利用しているシェルに環境変数を設定してください。
以下にfishの場合の例を示します。
```sh
set -x DB_PORT 5432
set -x CGO_ENABLED 0
set -x PASSWORD password
set -x DBNAME wantedly
set -x USER postgres
set -x HOSTNAME localhost
```

ローカルで開発する場合、`HOSTNAME`を`localhost`に設定してください。

### Goのコードをコンパイルする

このプロダクトは[Go modules](https://github.com/golang/go/wiki/Modules)を利用しているため`go build`でビルドを行うことができます。

有効にしていない場合、環境変数に`GO111MODULES=on`を設定してください。

