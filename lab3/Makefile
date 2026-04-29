
EXE_NAME=main

.PHONY: dev migrate

dev:
	@echo "Starting development server..."
	cd .\backend\cmd && go build -o $(EXE_NAME) main.go
	.\$(EXE_NAME)

migrate:
	docker-compose run --rm backend ./server