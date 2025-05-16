CUR_DIR := $(PWD)
ENV_VARS_PATH := --env-file "$(CUR_DIR)/.env"
DB_IMAGE := postgres:15-alpine
DB_VOLUME_PARAM := -v "$(CUR_DIR)/infra/postgres/scripts:/docker-entrypoint-initdb.d"

db/up:
	@make db/down
	@docker run --rm -d -p 5432:5432 \
		$(ENV_VARS_PATH) \
		$(DB_VOLUME_PARAM) \
		--name postgres $(DB_IMAGE) > /dev/null
db/down:
	@docker container stop postgres &> /dev/null || true

test:
	@ginkgo -r

test/integration: db/up test db/down

up:
	@docker-compose -f "infra/docker/compose.yml" up

down:
	@docker-compose -f "infra/docker/compose.yml" down

build:
	@docker-compose -f "infra/docker/compose.yml" build

start:
	@make down
	@make build
	@make up
