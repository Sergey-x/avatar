env:
	cp .env.example .env

startDb:
	docker compose -f docker-compose-dev.yml up

run:
	go build && ./avatar
