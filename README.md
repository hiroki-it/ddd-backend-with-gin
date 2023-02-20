# ddd-backend-with-gin

## 目次

<!-- TOC -->
* [概要](#概要)
* [ディレクトリ構成](#ディレクトリ構成)
* [環境構築](#環境構築)
  * [ビルドイメージ & コンテナ構築](#ビルドイメージ--コンテナ構築)
* [ホットリロード](#ホットリロード)
<!-- TOC -->

<br>

## 概要

Clean-Arch，Go，Ginを学習するためのアプリケーションです．

> ↪️ 参考：
> 
> - [クリーンアーキテクチャについて](https://hiroki-it.github.io/tech-notebook/software/software_application_architecture_backend_domain_driven_design_clean_architecture.html)

## ディレクトリ構成

本リポジトリのディレクトリ構造は以下の通りに構成しております．

クリーンアーキテクチャに則ったレイヤー名，パッケージ名，依存関係を取り入れております．

また，インターフェースは各レイヤーパッケージのルートに配置しました．

ただし，APIとして使用するため，インターフェース層のプレゼンター，ユースケース層のアウトプットバウンダリを廃止しております．

これに伴い，ユースケース層のインターラクターは，プレゼンターではなくレスポンスモデルを返却するようにしております．

実装方法は、[こちら](https://hiroki-it.github.io/tech-notebook/software/software_application_architecture_backend_domain_driven_design_clean_architecture.html)に整理しております．

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
    |   ├── database         # データベース
    |   ├── logger           # ロガー 
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
            ├── interactors # インターラクター
            ├── boundaries  # インプットバウンダリ 
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

<br>

## ホットリロード

ホスト側のソースコードを修正に合わせて，コンテナ側のアーティファクトを随時更新します．

また，ホットリロードの実行時に，合わせてソースコード整形と静的解析を実行します．

ツールとして，[air](https://github.com/cosmtrek/air) を使用いたしました．

```shell
$ docker-compose exec app air -c .air.toml
```
