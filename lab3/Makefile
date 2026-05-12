
EXE_NAME=main

.PHONY: dev migrate

dev:
	@echo "Starting development server via Docker Compose..."
	docker compose --env-file ./backend/cmd/boot.env up -d --build

migrate:
	docker compose --env-file ./backend/cmd/boot.env run --rm migrate

clean-db:
	@echo "Removing Docker containers and clearing database volumes..."
	docker compose --env-file ./backend/cmd/boot.env down -v