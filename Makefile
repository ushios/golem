
build: dep
	docker-compose build golem

test:
	docker-compose run --rm golem go test -v -cover ./...

dep:
	docker-compose run --rm dep
