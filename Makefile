# 環境変数読み込み
ifneq (,$(wildcard .env))
	include .env
	export 
endif

# DB
db = mysql -u $(MYSQL_USER) -p$(MYSQL_PASSWORD) -h 127.0.0.1 -P 3306 $(MYSQL_DATABASE)

# docker
up:
	@docker compose up -d
down:
	@docker compose down

migrate-up:
	$(db) < $(firstword $(wildcard migrations/up/*.sql))

migrate-down:
	$(db) < $(firstword $(wildcard migrations/down/*.sql))