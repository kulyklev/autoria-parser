SHELL := /bin/bash

# ======================================================================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all

check:
	go vet ./...
	staticcheck ./app/... ./internal/...

dump:
	docker exec -t parser-postgres pg_dumpall -c -U postgres > dump_`date +%d-%m-%Y"_"%H_%M_%S`.sql

restore:
	cat your_dump.sql | docker exec -i parser-postgres psql -U postgres


run:
	go run gitlab.com/kulyklev/autoria-parser
