migrate:
	@go run cmd/migrate/main.go migrate

seed:
	@go run cmd/migrate/main.go seed