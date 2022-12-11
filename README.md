### Setup with docker

- `cp .env.example .env` and update variables
- `docker-compose -f docker-compose.yml up -d --build`

### Setup without using docker

- `make install`
- `make start`

### Run migration

- `make migrate`
- `make migrate-rollback`
- `make create-migration create_users_table`

### Production setup

- `docker-compose -f docker-compose.production.yml up -d --build`
