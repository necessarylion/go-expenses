### Setup with docker

- `cp .env.example .env` and update variables
- `docker-compose -f docker-compose.yml up -d --build`

### Setup without using docker

- `go install github.com/cosmtrek/air@latest` (Live reload for Go apps)
- `go install github.com/Altoros/gorm-goose/cmd/gorm-goose@latest` (Migration tool)
- `cp .env.example .env` and update variables
- `go mod tidy`
- `make run`

### Run migration

- `go get github.com/Altoros/gorm-goose/lib/gorm-goose` (Optional If failed to run migration)
- `make migrate`
- `make migrate-rollback`
- `make create-migration create_users_table`

### Production setup

- `docker-compose -f docker-compose.production.yml up -d --build`
