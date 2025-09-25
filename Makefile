# Makefile for the Severance project

# ====================================================================================
# VARIABLES
# ====================================================================================
DB_USER=root
DB_PASSWORD=secret
DB_NAME=severance
DB_PORT=5432
# This is the connection string for the migrate tool.
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable
# This is our Go module path from go.mod
MODULE_PATH=github.com/your-github-username/severance

# ====================================================================================
# DATABASE & DOCKER COMMANDS
# ====================================================================================
postgres:
	# Run a new postgres docker container. We name it severance-db for easy access.
	# We also check if it's already running to avoid errors.
	docker run --name severance-db -p $(DB_PORT):5432 -e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) -d postgres

createdb:
	# Create the application database inside the docker container.
	docker exec -it severance-db createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

dropdb:
	# Drop the application database. Useful for a clean reset.
	docker exec -it severance-db dropdb $(DB_NAME)

# ====================================================================================
# MIGRATIONS
# ====================================================================================
# We'll create the db/migration folder in the next step.
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

# ====================================================================================
# CODE GENERATION & TESTING
# ====================================================================================
sqlc:
	# Generate type-safe Go code from our SQL queries.
	# This command depends on a sqlc.yaml file which we will create soon.
	sqlc generate

test:
	# Run all tests verbosely and with coverage.
	go test -v -cover ./...

# The mock target is for later when we write unit tests and need to mock the database.
# mock:
# 	mockgen -package mockdb -destination internal/database/mock/store.go $(MODULE_PATH)/internal/database/sqlc Store

# ====================================================================================
# APPLICATION COMMANDS
# ====================================================================================
server:
	# Run the application. Note the path is now correct for our project structure.
	go run ./cmd/severance/main.go

# ====================================================================================
# TOOLING INSTALLATION
# ====================================================================================
install-tools:
	# Install the command-line tools we need.
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	# go install go.uber.org/mock/mockgen@latest


# .PHONY ensures that make doesn't confuse these targets with actual files.
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server install-tools
