SHELL := /bin/bash

include .env
export $(shell sed 's/=.*//' .env)

# generate-env:
# 	@./generate_db_url.sh > .generated_env
# 	@echo "Generated .generated_env:"
# 	@cat .generated_env
# 	@source .generated_env

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'initial_migration'
.PHONY: migrate-create

migrate-up:  ### migration up
	@source .generated_env && migrate -path migrations -database "$$DB_URL" up
.PHONY: migrate-up

migrate-down:  ### migration down
	@source .generated_env && migrate -path migrations -database "$$DB_URL" down
.PHONY: migrate-down

migrate-force:  ### force migration to version 1
	@source .generated_env && migrate -path migrations -database "$$DB_URL" force 1
.PHONY: migrate-force

migrate-file:  ### create a new migration file
	@source .generated_env && migrate create -ext sql -dir migrations/ -seq init_tables
.PHONY: migrate-file

proto:
	./scripts/gen-proto.sh


give-permissions:
	- chmod +x additional.sh
	- chmod +x generate_db_url.sh


.PHONY: migrate run

.PHONY: migrate run

psql:
	psql -U postgres -d postgres

run: proto
	go run cmd/main.go

.PHONY: run, proto
