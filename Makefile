start:
	@clear
	@echo "\n Start Docker ... \n"
	@docker-compose up -d
	@clear
	@echo "\n Done! \n"

stop:
	@clear
	@echo "\n Stop Docker ... \n"
	@docker-compose down
	@clear
	@echo "\n Done! \n"

run:
	@clear
	@echo "\n  Service Online <Prod> ... \n"
	@go build -o ./bin/cmd ./cmd/main.go
	@./bin/cmd

dev:
	@clear
	@echo "\n  Service Online <Dev> ...  \n"
	@go run ./cmd/main.go

test:
	@clear
	@echo "\n  Testing ...  \n"
	@go test -v ./...


build:
	@clear
	@echo "\n  Building ...  \n"
	@rm -rf ./bin
	@go build -o ./bin/cmd ./cmd/main.go