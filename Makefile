start:
	docker-compose up \
		--detach \
		--build \
		--remove-orphans \
		--force-recreate

stop:
	docker-compose stop

logs:
	docker-compose logs --follow