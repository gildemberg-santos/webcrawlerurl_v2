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
	@go build -o ./bin/app ./cmd/main.go
	@./bin/app

dev:
	@clear
	@echo "\n  Service Online <Dev> ...  \n"
	@go run ./cmd/main.go
