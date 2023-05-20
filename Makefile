createmigration:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/fin_api_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/fin_api_db?sslmode=disable" -verbose down

.PHONY: migrateup migratedown createmigration
