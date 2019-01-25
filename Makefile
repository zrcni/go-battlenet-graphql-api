start:
	go run ./server/server.go

generate:
	go run ./scripts/gqlgen.go

build:
	go build -o ./bin/api ./server/server.go

buildto:
	@if [ -z $(path) ]; then\
		echo "path arg is empty. Usage: make $@ path=/path/to/file";\
		exit 1;\
	fi
	CGO_DISABELD=0 GOOS=linux GOARCH=amd64 go build -o $(path) ./server/server.go
