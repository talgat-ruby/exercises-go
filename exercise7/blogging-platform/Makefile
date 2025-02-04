ifneq (,$(wildcard .env))
    include .env
    export
endif

up:
	migrate -database $(POSTGRES_URL) -path db/migrations/ up

down:
	migrate -database $(POSTGRES_URL) -path db/migrations/ down