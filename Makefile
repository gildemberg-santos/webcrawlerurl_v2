start:
	@echo "Executando o conteiner..."
	@docker-compose up -d
	@echo "Fim da execução do conteiner!"
	@make run

stop:
	@echo "Parando o conteiner..."
	@docker-compose down
	@echo "Fim da execução do conteiner!"

run:
	@clear
	@echo "Executando ..."
	@go build -o ./cmd/main ./cmd/main.go
	@go run ./cmd/main.go
	@echo "Fim da execução!"
