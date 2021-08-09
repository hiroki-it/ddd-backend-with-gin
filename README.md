# ddd-api-with-go-gin

## ディレクトリ構成

クリーンアーキテクチャに則ったデザインパターンと依存関係を取り入れております．

一点，APIとして使用するため，プレセンターを廃止しております．

これに伴い，ユースケース層のインターラクターはプレゼンターではなく，レスポンスモデルを返却するようにしております．

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
    ├── infrastructure       # ** インフラストラクチャ層 **
    |   ├── middlewares      # ミドルウェア
    |   ├── routers          # ルータ
    |   └── foo              # 任意のルートエンティティ
    |       ├── dtos         # gormモデルからドメインモデルへの変換
    |       ├── repositories # 実装リポジトリ
    |       └── seeders      # DBテストデータ
    |
    ├── interfaces          # ** インターフェース層 **
    |   └── foo             # 任意のルートエンティティ
    |       └── controllers # コントローラ
    |     
    └── usecase             # ** ユースケース層 **
        └── foo             # 任意のルートエンティティ
            ├── inputs      # インプットポート 
            ├── interactors # インターラクタ 
            ├── requests    # リクエストモデル
            └── responses   # レスポンスモデル
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
$ docker-compose run --rm app go mod download -x
```

<br>

## ホットリロード

ホスト側のソースコードを修正に合わせて，コンテナ側のアーティファクトを随時更新します．

また，ホットリロードの実行時に，合わせてソースコード整形と静的解析を実行します．

ツールとして，[air](https://github.com/cosmtrek/air) を使用いたしました．

```shell
$ docker-compose run --rm app air -c .air.toml
```
