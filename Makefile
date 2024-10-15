DB_DSN := "postgres://postgres:admin@localhost:5433/restapi_dev?sslmode=disable"

MIGRATE := migrate -database $(DB_DSN) -path ./migrations 

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/app/main.go

gen:
	oapi-codegen -package messages -config openapi/.openapi -include-tags messages  openapi/openapi.yaml > ./internal/web/messages/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number