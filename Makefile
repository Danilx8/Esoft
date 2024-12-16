up: rebuild
	docker compose up -d

up-db:
	docker compose up db -d

rebuild: down
	docker compose build

down:
	docker compose down

migrate: up-db
	for file in $$(ls sql_scripts/); do \
		docker exec -t esoft-db-1 psql -U example -d polytech-esoft -f "/sql_scripts/$$file"; \
	done

logs:
	docker compose logs -f
