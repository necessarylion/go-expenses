include .env
export

install:
	go install github.com/cosmtrek/air@latest
	go install github.com/Altoros/gorm-goose/cmd/gorm-goose@latest
	go mod tidy
	@[ ! -f ./.env ] && cp .env.example .env || true

start:
	air

migrate:
	gorm-goose up

migrate-rollback:
	gorm-goose down

create-migration:
	gorm-goose create $(filter-out $@,$(MAKECMDGOALS))

%:
	@: