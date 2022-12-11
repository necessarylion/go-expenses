include .env
export

run:
	air

migrate:
	gorm-goose up

migrate-rollback:
	gorm-goose down

create-migration:
	gorm-goose create $(filter-out $@,$(MAKECMDGOALS))

%:
	@: