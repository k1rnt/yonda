# yonda
本の進捗状況を管理

# setup
1. `docker compose up -d`
2. `cp .env.example .env`

# migrate
golang-migrateが入ってたら1はスキップしておk
1. `brew install golang-migrate`
2. `migrate -source file://database/migrate -database 'mysql://k1rnt:k1rnt_pass@tcp(localhost:3306)/yonda' up`

# api run
1. `cd api`
2. `go run cmd/yonda/*.go`

# golangci-lint
golangci-lintが入ってたら1はスキップしておk  
1. `brew install golangci-lint`
2. `cd api`
3. `golangci-lint run --config=golangci.yml`
