# go-JWT
2025.09 ~ JWT実装用

環境変数のコピー
```
cp .env.sample .env
```

## コンテナ操作
```
# コンテナ立ち上げ
make up

# コンテナ停止
make down
```

## マイグレーションの実行
```
# up
make migrate-up

# down
make migrate-down
```