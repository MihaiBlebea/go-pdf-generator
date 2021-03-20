setup: env-file build up

setup-test: env-file build-test up-test cover-html

env-file: 
	cp ./.env.example ./.env

build:
	docker build \
		-t serbanblebea/go-diploma:v0.1 \
		.

up:
	docker run \
		--rm \
		--name go-diploma \
		-p 8087:8087 \
		-v ${PWD}/storage:/app/storage \
		--env-file ./.env \
		serbanblebea/go-diploma:v0.1

stop: 
	docker stop go-diploma

build-test:
	docker build \
		--no-cache \
		--file ./Dockerfile.test \
		-t serbanblebea/go-diploma:test \
		.

up-test:
	docker run \
		-v ${PWD}:/app \
		--rm \
		--name go-diploma-test \
		--env-file ./.env \
		serbanblebea/go-diploma:test

cover-html:
	go tool \
		cover -html=cover.out \
		-o cover.html \
		&& open cover.html

go-build:
	go build -o=./diploma .

go-test:
	go test -v ./...

generate:
	docker exec -it go-diploma /bin/sh -c "./diploma generate -c 1000"
