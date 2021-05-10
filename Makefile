RUN:
	go run ./cmd/api

CREATE_MOVIES_TABLE_MIGRATION:
	migrate create -seq -ext=.sql -dir=./migrations create_movies_table

ADD_MOVIES_CHECK_CONSTRAINTS_MIGRATION:
	migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constraints

EXECUTE_UP_MIGRATIONS:
	migrate -path=./migrations -database=$$GREENLIGHT_DB_DSN up

EXECUTE_DOWN_MIGRATIONS:
	migrate -path=./migrations -database=$$GREENLIGHT_DB_DSN down

ADD_MOVIES_INDEXES_MIGRATIONS:
	migrate create -seq -ext .sql -dir=./migrations add_movies_indexes