# gakujo-notification

未提出課題を一覧表示するアプリです。一応クロスプラットフォームですが、サーバーのデプロイとかがまだできていません。

## スクリーンショット

<img width="300" alt="Screen Shot 2022-07-30 at 12 46 31" src="https://user-images.githubusercontent.com/43411965/181871245-4e8c4932-2c2c-442a-a05c-46aefea4b0f4.png">
<img width="300" alt="Screen Shot 2022-07-30 at 12 46 49" src="https://user-images.githubusercontent.com/43411965/181871255-5d83d92e-7f5a-4dfa-95bd-3fa031376fb5.png">
<img width="300" src=https://user-images.githubusercontent.com/43411965/181871941-f4b13f81-f310-4ad5-bfb1-e9bb831e24d4.gif />

## ビルド手順

### Requirements

括弧内は開発環境でのバージョンです

- Docker(20.10.17): https://www.docker.com/get-started/
- nodejs(v16.16.0): https://nodejs.org/ja/
- npm(8.11.0): https://www.npmjs.com/
- yarn(1.22.19): https://yarnpkg.com/
- expo-go(latest): https://expo.dev/client (on android or ios)

### Steps

1. バックエンドを起動する

```shell
$ pwd
/path/to/gakujo-notification

$ docker compose up -d --build
.
.
[+] Running 2/2
 ⠿ Container gakujo-notification-db-1       Running                                                  0.0s
 ⠿ Container gakujo-notification-backend-1  Started                                                  0.3s
```

2. `gakujo-notification-frontend/.env` を作成する

自分の PC のネットワーク上の IP アドレスを調べて `API_URL` にセットしてください。

```
API_URL=http://<network-ip-addr>:8080
```

3. モバイルアプリのビルド

以下のように `$ yarn start` を実行すると以下のような qr コードが表示されるので、expo-go をインストールしてある端末の qr コードリーダーでかざしてください。かざしてしばらく待つとアプリが起動すると思います。

> **Note**
> パソコンと端末は同一のネットワークに接続してください。そうしないとパソコン側で起動している api サーバーに端末から接続することができません。

```shell
$ cd gakujo-notification-frontend
$ yarn install   # 依存関係のインストール
$ yarn start
▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
█ ▄▄▄▄▄ █▀▀ ████  █ ▄▄▄▄▄ █
█ █   █ █▄▀██▀█▄▄██ █   █ █
█ █▄▄▄█ █ ▄ █  ▀ ██ █▄▄▄█ █
█▄▄▄▄▄▄▄█ █ ▀▄█▄▀ █▄▄▄▄▄▄▄█
█  ▀▄▀█▄ █▀█   ▀██▀  ▄▀▄▄▀█
██▄█▄  ▄▄ ▀▀    ▀▄▄▀ ▀▀█▄▄█
███▄   ▄▄██▄ ▄  █▀█ ▄█ ██▀█
█▄▀▄█▀ ▄█ ▀█ ▄ ██ ▄▄ ▀▀██▄█
█▄▄█▄██▄▄  ▀▀█▀▄  ▄▄▄ █ ▄ █
█ ▄▄▄▄▄ █▄█▀▀▄█   █▄█  ▀  █
█ █   █ █▀▀▀▄█▄▀▀▄ ▄▄ █▀▄██
█ █▄▄▄█ █▀  █▄ ▄█  █▄  ▄█▄█
█▄▄▄▄▄▄▄█▄▄███▄▄█▄███▄▄█▄▄█
```

## プロジェクト(ソースコード)の説明

**gakujo-notification-backend**

一部省略しています

```
.
├── main.go            # エントリーポイント
├── gakujo/            # 学情のスクレイピングを行うライブラリ群。モデルとかも全てここで定義済     
│  ├── assignments.go  # 課題一覧をヘッドレスブラウザでクローリングする
│  ├── client.go       # Client 構造体や utility 関数の定義
│  ├── enum.go         # AssignmentKind や Status などの enum の定義
│  ├── scrape.go       # クローリングで取得した html を対象にスクレイピングする
│  ├── scrape_test.go
│  ├── web-crawler.go  # ヘッドレスブラウザでログインしたりする関数の定義
│  └── web-crawler_test.go
├── repository/
│  ├── assignment.go   # 課題一覧のテーブルやそれに対するクエリを定義
│  ├── repository.go   # Client 構造体を定義
│  ├── repository_test.go
│  ├── user.go         # ユーザ一覧のテーブルやそれに対するクエリを定義
│  └── user_assignment.go  # UserAssignment テーブルやそれに対するクエリを定義(ER図を参照)
└── server/            # web-api のハンドラやルーティングを定義
```

**ER図**

<img width="500" alt="Screen Shot 2022-07-30 at 13 06 40" src="https://user-images.githubusercontent.com/43411965/181871744-9f21fd2f-b8b1-4f6b-85b4-cfe530af1f40.png">
