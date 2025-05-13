# Load environment variables from .env if available
ifneq ("$(wildcard .env)","")
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# --- Configuration ---
POSTGRES_DSN ?= $(shell printenv POSTGRES_DSN)
POSTGRES_TEST_DSN ?= $(shell printenv POSTGRES_TEST_DSN)

# --- Helpers ---
check-db-env:
	@test -n "$(POSTGRES_DSN)" || (echo "❌ POSTGRES_DSN is not set" && exit 1)

check-test-db-env:
	@test -n "${POSTGRES_TEST_DSN}" || (echo "❌ POSTGRES_TEST_DSN is not set" && exit 1)

GET_FIRST_USER = $(shell psql $(POSTGRES_DSN) -At -c "SELECT id FROM users LIMIT 1;")
GET_USER_IDS = $(shell psql $(POSTGRES_DSN) -At -c "SELECT id FROM users;")

GET_FIRST_ROADMAP = $(shell psql $(POSTGRES_DSN) -At -c "SELECT id FROM roadmaps LIMIT 1;")
GET_ROADMAP_IDS = $(shell psql $(POSTGRES_DSN) -At -c "SELECT id FROM roadmaps;")

GET_FIRST_CHECKPOINT = $(shell psql $(POSTGRES_DSN) -At -c "SELECT id FROM checkpoints LIMIT 1;")
GET_CHECKPOINT_IDS = $(shell psql $(POSTGRES_DSN) -At -c "SELECT id FROM checkpoints;")

# --- Dev Commands ---
run:
	go run cmd/server/main.go

# --- Migrations ---
migrate-up: check-db-env
	migrate -path migrations -database "$(POSTGRES_DSN)" up

migrate-down: check-db-env
	migrate -path migrations -database "$(POSTGRES_DSN)" down

migrate-test-up: check-test-db-env
	migrate -path migrations -database "$(POSTGRES_TEST_DSN)" up 

migrate-test-down: check-test-db-env
	migrate -path migrations -database "$(POSTGRES_TEST_DSN)" down

# --- Database Seeding --- 
db-seed: check-db-env
	@echo "☘️ Seeding users, roadmaps, checkpoints"
	psql $(POSTGRES_DSN) -f seed/users.sql
	psql $(POSTGRES_DSN) -f seed/roadmap.sql
	psql $(POSTGRES_DSN) -f seed/checkpoints.sql
	@echo "✅ Seed complete"

# --- gRPC Testing ---
first-user: check-db-env
	$(eval USER_ID := $(shell echo $(GET_FIRST_USER) | tr -d '\n'))
	@echo "⏳ Retrieving first user: $(USER_ID)"
	grpcurl -plaintext -d '{"id":"$(USER_ID)"}' localhost:9090 proto.users.v1.UserService/GetUser;

fetch-all-users: check-db-env
	@echo "🧪 Testing all users:"
	@for id in $(GET_USER_IDS); do \
		if [ $$(echo $$id | wc -c) -eq 37 ]; then \
			echo "⏳ Retrieving user ID: $$id"; \
			grpcurl -plaintext -d "{\"id\":\"$$id\"}" localhost:9090 proto.users.v1.UserService/GetUser; \
		else \
			echo "⚠️ Skipping invalid UUID: $$id"; \
		fi; \
	done
	
first-roadmap: check-db-env
	$(eval ROADMAP_ID := $(shell echo $(GET_FIRST_ROADMAP) | tr -d '\n'))
	@echo "⏳ Retrieving roadmap ID: $(ROADMAP_ID)"
	grpcurl -plaintext -d '{"roadmap_id":"$(ROADMAP_ID)"}' localhost:9090 proto.roadmaps.v1.RoadmapService/GetRoadmap;

fetch-all-roadmaps: check-db-env
	@echo "⏳ Retrieving all roadmaps:"
	@for id in $(GET_ROADMAP_IDS); do \
		if [ $$(echo $$id | wc -c ) -eq 37 ]; then \
			echo "🔸 Retrieving roadmap ID: $$id"; \
			grpcurl -plaintext -d "{\"roadmap_id\":\"$$id\"}" localhost:9090 proto.roadmaps.v1.RoadmapService/GetRoadmap; \
		else \
			echo "⚠️ Skipping invalid UUID: $$id"; \
		fi; \
	done 

first-checkpoint: check-db-env
	$(eval CHECKPOINT_ID := $(shell echo $(GET_FIRST_CHECKPOINT) | tr -d '\n'))
	@echo "⏳ Retrieving first checkpoint ID: $(CHECKPOINT_ID)"
	grpcurl -plaintext -d '{"checkpoint_id":"$(CHECKPOINT_ID)"}' localhost:9090 proto.checkpoints.v1.CheckpointService/GetCheckpoint;

fetch-all-checkpoints: check-db-env
	@echo "⏳ Retrieving all checkpoints"
	@for id in $(GET_CHECKPOINT_IDS); do \
		if [ $$(echo $$id | wc -c) -eq 37 ]; then \
			echo "🔸 Retrieving checkpoint ID: $$id"; \
			grpcurl -plaintext -d "{\"checkpoint_id\":\"$$id\"}" localhost:9090 proto.checkpoints.v1.CheckpointService/GetCheckpoint; \
		else \
			echo "⚠️ Skipping invalid UUID: $$id"; \
		fi; \
	done

# --- Tests ---
test: ## Run all tests for users
	go test -v ./tests/...

test-users: 
	go test -v ./tests/users/...

test-checkpoints: 
	go test -v ./tests/checkpoints/...

test-roadmaps:
	go test -v ./tests/roadmap/...

# --- Meta ---

.PHONY: run migrate-up migrate-down db-seed \
	migrate-up-test-users migrate-down-test-users \
	test test-users test-checkpoints test-roadmaps \
	check-db-env check-test-db-env \
	db-seed first-user all-users \
	first-roadmap fetch-all-roadmaps first-checkpoint fetch-all-checkpoints
