
build: dep
	docker-compose build golem-build

test:
	docker-compose run --rm golem-build go test -v -cover ./...

dep:
	docker-compose run --rm dep

docker:
	docker build -t golem .

binary: _binary_linux _binary_darwin _binary_windows

_binary_linux:
	docker-compose run --rm golem-build env GOOS=linux GOARCH=amd64 go build  -o ./bin/linux/amd64/golem cmd/golem/main.go

_binary_darwin:
	docker-compose run --rm golem-build env GOOS=darwin GOARCH=amd64 go build  -o ./bin/darwin/amd64/golem cmd/golem/main.go

_binary_windows:
	docker-compose run --rm golem-build env GOOS=windows GOARCH=amd64 go build  -o ./bin/windows/amd64/golem cmd/golem/main.go
