up: rebuild
	docker compose up -d

up-db:
	docker compose up db -d

rebuild: down swagger up-db
	docker compose build

down:
	docker compose down

migrate: up-db
	for file in $$(ls -tr sql_scripts/); do \
		docker exec -t esoft-db-1 psql -U example -d polytech-esoft -f "/sql_scripts/$$file"; \
	done

swagger:
	swag init -g cmd/main.go -d app/ --pd true -q

logs:
	docker compose logs -f
