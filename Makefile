create-migration:
	atlas migrate diff --env gorm

migrate:
	@go run cmd/migrate/main.go migrate

seed:
	@go run cmd/migrate/main.go seed

compute-migrate-checksum:
	atlas migrate hash