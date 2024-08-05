back = api
front = front


run: 
	cd $(back)/cmd && go run main.go

db-up: 
	cd $(back)/docker && docker compose up

db-down:
	cd $(back)/docker && docker compose down


dev:
	cd $(front) && pnpm run dev

db-seed: 
	cd $(back)/internal/database && go run scripts/seed_script.go