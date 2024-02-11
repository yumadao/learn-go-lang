# start db
start db:
	docker compose up -d

# remove db
remove db:
	docker compose rm -s -f -v

# start app
start:
	GO_ENV=dev go run .

# run migrate
migrate:
	GO_ENV=dev go run migrate/migrate.go