
build: dep
	docker-compose build golem-build

test:
	docker-compose run --rm golem-build go test -v -cover ./...

dep:
	docker-compose run --rm dep
