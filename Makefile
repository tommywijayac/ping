-include .env

.PHONY: build

# build
build:
	@echo " > Building [ping]..."
	@cd ./cmd/ping/ && go build -o ../../bin && cd ../..
	@echo " > Finished building [ping]"

run: build
	@echo " > Running [ping]..."
	@./bin/ping
	@echo " > Finished running [ping]"
