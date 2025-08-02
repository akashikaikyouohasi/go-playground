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

# Go モジュール関連
.PHONY: mod-init mod-tidy mod-clean
mod-init:
    docker-compose exec go go mod init go-playground # モジュールの依存関係を管理するファイルを初期化
mod-tidy:
    docker-compose exec go go mod tidy # 不要な依存関係を削除し、必要な依存関係を追加
mod-clean:
    docker-compose exec go go clean -modcache

# テスト関連
test:
	go test
test-all:
	docker-compose exec go go test -v ./... -coverprofile=coverage.out
	docker-compose exec go go tool cover -html=coverage.out -o coverage.html
