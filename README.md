# ddd-api-with-go-gin

## ディレクトリ構成

```
project
├── build # ビルド処理
└── cmd   # エントリポイントを含む処理，ユニットテスト
    ├── domain               # ** ドメイン層 **
    |   └── foo              # 任意のルートエンティティ
    |       ├── entities     # エンティティ
    |       ├── ids          # ID
    |       ├── repositories # インターフェースリポジトリ
    |       └── values       # 値オブジェクト
    |
    ├── infrastructure       # ** インターフェース層 **
    |   ├── middlewares      # ミドルウェア
    |   ├── routers          # ルータ
    |   └── foo              # 任意のルートエンティティ
    |       ├── dtos         # gormモデルからドメインモデルへの変換
    |       ├── repositories # 実装リポジトリ
    |       └── seeders      # DBテストデータ
    |
    ├── interfaces          # ** インターフェース層 **
    |   └── foo             # 任意のルートエンティティ
    |       ├── controllers # コントローラ
    |       └── presenters  # プレゼンター（レスポンスデータの構築）
    |     
    └── usecase             # ** ユースケース **
        └── foo             # 任意のルートエンティティ
            ├── inputs      # インプット（インターラクターのパラメータの構築 & バリデーション）
            └── interactors # インターラクター
```

## 環境構築

### ビルドイメージ & コンテナ構築

Dockerfileからイメージをビルドし，コンテナを構築します．

イメージのビルドにおいては，マルチステージビルドを採用しております．

```shell
$ docker-compose run -d
````

### モジュールのインストール

コンテナで，アプリケーションで使用するモジュールをインストールします．

```shell
$ docker-compose run --rm api go mod download -x
```

<br>

## ホットリロード

ホスト側のソースコードを修正に合わせて，コンテナ側のアーティファクトを随時更新します．

また，ホットリロードの実行時に，合わせてソースコード整形と静的解析を実行します．

ツールとして，[air](https://github.com/cosmtrek/air) を使用いたしました．

```shell
$ docker-compose run --rm api air -c .air.toml
```
