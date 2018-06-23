
build: dep
	docker-compose build build-golem

test:
	docker-compose run --rm build-golem go test -v -cover ./...

dep:
	docker-compose run --rm dep
