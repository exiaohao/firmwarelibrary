DATETIME = $(shell date +'%Y%m%d%H%M%S')

image:
	GOOS=linux GOARCH=amd64 go build -o ./docker/flareproxy ./cmd/server

