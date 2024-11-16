
GO = env GOOS=linux GOARCH=amd64 go

.PHONY: clean build_app build run stop restart logs

build: clean build_app

restart: build stop run

run:
	docker stop articles_api || true
	docker rm articles_api || true
	docker run -ti \
    	--network dev \
    	--name articles_api \
    	-w /usr/src/app \
    	-v "$(PWD)":/usr/src/app \
    	alpine:3.14 build/app

stop:
	-docker stop articles_api

restart: stop build run

clean:
	@echo "Clean.."
	rm build/* || true

build_app:
	@echo "Building..."
	$(GO) build -o build/app cmd/http/main.go

logs:
	docker logs -f articles_api
