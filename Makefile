.PHONY: build
build:
	go build -o ./bin/obs main.go

.PHONY: server
server:
	air -c .air.toml

.PHONY: test
test:
	go test -v ./...

.PHONY: test.watch
test.watch:
	while inotifywait -e close_write ./pkg/*; do go test ./... -json | tparse -all; done

.PHONY: db.migrate
db.migrate:
	migrate -path ./migrations -database $(DATABASE_URL) -verbose up

.PHONY: db.reset
db.reset:
	migrate -path ./migrations -database $(DATABASE_URL) -verbose drop -f
	migrate -path ./migrations -database $(DATABASE_URL) -verbose up

.PHONY: db.migration.create
db.migration.create:
	migrate create -ext sql -dir migrations -seq ${name}

.PHONY: db.psql
db.psql:
	docker exec -it -u postgres rimshot-db-1 psql -d obs_dev