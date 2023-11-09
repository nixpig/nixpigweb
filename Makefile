dev:
	air --build.cmd "go build -o bin/api cmd/api/main.go" --build.bin "./bin/api"

run:
	go run cmd/api/main.go

test:
	go test -v cmd/api/./..

clean:
	rm -rf bin tmp cmd/api/tmp

build:
	go build -o bin/api cmd/api/main.go

