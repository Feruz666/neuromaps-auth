postgres:
	docker run --name=nmauth-db-hub -e POSTGRES_PASSWORD='qwerty' -p 5435:5432 -d --rm umberman/postgres

migrateup:
	migrate -path ./schema -database "postgresql://postgres:qwerty@localhost:5435/postgres?sslmode=disable" -verbose up

run:
	go run main.go

.PHONY:
	migrateup postgres run