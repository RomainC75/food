testsum:
	gotestsum --watch --format testname -- -v ./... 

migrateup:
	migrate --path db/migration --database "postgresql://name:pass@localhost:5432/securitest?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://name:pass@localhost:5432/securitest?sslmode=disable" --verbose down

sqlc:
	sqlc generate

.PHONY: testsum migrateup migratedown sqlc 