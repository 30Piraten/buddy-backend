# Load environment variables from .env if available
ifneq ("$(wildcard .env)","")
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# Database variables
POSTGRES_DSN ?= $(shell printenv POSTGRES_DSN)

# Helper to check if POSTGRES_DSN is set
check-db-env:
	@test -n "$(POSTGRES_DSN)" || (echo "POSTGRES_DSN is not set" && exit 1)

# Dynamic user ID helpers
GET_FIRST_USER = $(shell psql $(POSTGRES_DSN) -t -c "SELECT id FROM users LIMIT 1;")
GET_USER_IDS = $(shell psql $(POSTGRES_DSN) -t -c "SELECT id FROM users;")

# Run the application
run:
	go run cmd/server/main.go

# Migrate database (up or down)
migrate-up: check-db-env
	migrate -path migrations -database "$$POSTGRES_DSN" up

migrate-down: check-db-env
	migrate -path migrations -database "$$POSTGRES_DSN" down

migrate-new:
	migrate create -ext sql -dir migrations -seq

# Seed the database
db-seed: check-db-env
	psql $(POSTGRES_DSN) -f scripts/seed.sql

# Test with the first seeded user
test-first-user: check-db-env
	$(eval USER_ID := $(shell echo $(GET_FIRST_USER) | tr -d '\n' | xargs))
	@echo "Testing with first user: $(USER_ID)"
	grpcurl -plaintext -d '{"id":"$(USER_ID)"}' localhost:9090 users.v1.UserService/GetUser

# Test with all users
test-all-users: check-db-env
	@for id in $(shell echo $(GET_USER_IDS) | tr -d '[:space]' | sed 's/|/ /g'); do \
		if [ ! -z "$$id" ]; then \
			echo "Testing user ID: $$id"; \
			grpcurl -plaintext -d "{\"id\":\"$$id\"}" localhost:9090 users.v1.UserService/GetUser; \
		fi; \
	done

.PHONY: run migrate-up migrate-down migrate-new db-seed check-db-env test-first-user test-all-users
