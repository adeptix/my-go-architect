create_db:
	docker-compose up -d

regenerate_models:
	sqlboiler psql

migrate:
	go run -v ./cmd/migration

fresh:
	go run -v ./cmd/migration -fresh

# show open ports
# sudo lsof -i -n -P | grep -E "\:5[0-9]{3}[^0-9]"

# generate migrations
# migrator create -d internal/migrations create_likes_table

