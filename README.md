## go-JWT

2025.09 ~ JWT実装用  
Go + Gin + GORM + MySQL で JWT認証付きAPIを実装

### ディレクトリ構成
- `main.go` ... エントリーポイント
- `db/` ... DB接続・リポジトリ
- `handler/` ... 各種APIハンドラー
- `middleware/` ... 認証ミドルウェア

---

### 主なAPI

- `POST /register` ... ユーザー登録
- `POST /login` ... ログイン（JWT発行）
- `GET /dashboard` ... 認証必須（JWTが必要）

---

### JWT認証について

- ログイン成功時、JWTトークンがCookieにセット
- 認証が必要なAPIは `middleware.Auth` で保護

---


### セットアップ

1. **環境変数ファイルの作成**

```bash
cp .env.sample .env
```

2. **Dockerで起動**


```bash
# コンテナ立ち上げ
make up

# コンテナ停止
make down
```

3. **マイグレーション**

```bash
# up
make migrate-up

# down
make migrate-down
```