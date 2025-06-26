# Dockerfileに更新があった場合、ビルドし直す用
build:
	docker-compose build --no-cache
	docker-compose up -d
up:
	docker-compose up -d
stop:
	docker-compose stop
down:
	docker-compose down
clean:
	docker-compose down --rmi all --volumes --remove-orphans
ps:
	docker-compose ps
logs:
	docker-compose logs
shell:
	docker-compose exec go bash

