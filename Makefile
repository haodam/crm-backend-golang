# These are the default values
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING = "root:root@tcp(127.0.0.1:3308)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema

up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

run:
	go run .\cmd\main.go

sql:
	sqlc generate


.PHONY: run up down reset sql