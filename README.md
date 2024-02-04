# learn-go-lang

setup db
docker compose up -d

Migration
GO_ENV=dev go run migrate/migrate.go

setup local server
GO_ENV=dev go run main.go
