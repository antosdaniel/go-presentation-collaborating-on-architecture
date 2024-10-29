MAKEFLAGS := --silent --always-make

# -r=false makes termination signals work properly
GOW := gow -r=false

.PHONY: dev
dev: install run-watch

.PHONY: run-watch
run-watch:
	$(GOW) -e go,mod -g=make run

.PHONY: run
run:
	go run .

.PHONY: generate
generate:
	go mod tidy && go mod vendor

.PHONY: install
install:
	go install github.com/mitranim/gow@latest
