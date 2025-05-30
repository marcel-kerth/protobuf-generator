compose:
	docker compose -f ./docker/prod/compose.yml up

compose_dev:
	docker compose -f ./docker/dev/compose.yml up