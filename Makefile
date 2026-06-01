DB_HOST=1###
DB_PORT=###
DB_USER=###
DB_PASS=###
DB_NAME=###

createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate: 
	migrate -path=sql/migrations \
	-database "mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" \
	-verbose up

migratedown:
	migrate -path=sql/migrations \
	-database "mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" \
	-verbose down

.PHONY: migrate migratedown createmigration