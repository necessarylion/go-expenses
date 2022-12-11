### Requirements

- `go install github.com/cosmtrek/air@latest` (Live reload for Go apps)
- `go install github.com/Altoros/gorm-goose/cmd/gorm-goose@latest` (Migration tool)
- `go get github.com/Altoros/gorm-goose/lib/gorm-goose` (Migration tool)

### Run migration

- `make migrate`
- `make migrate-rollback`
- `make create-migration create_users_table`
