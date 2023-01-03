LOCAL_DB_NAME:=orders_local
LOCAL_DB_DSN:=host=127.0.0.1 port=54320 dbname=orders_local user=orders_user password=orders_password sslmode=disable

db-reset-local:
	psql -c "drop database if exists $(LOCAL_DB_NAME)"
	psql -c "create database $(LOCAL_DB_NAME)"
	goose -dir internal/db/migrations postgres "$(LOCAL_DB_DSN)" up

db-create-migration:
	goose -dir internal/db/migrations create name sql

db-migrate:
	goose -dir internal/db/migrations postgres "$(LOCAL_DB_DSN)" up

db-migrate-down:
	goose -dir internal/db/migrations postgres "$(LOCAL_DB_DSN)" down
