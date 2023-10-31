prod:
	make test
	make run

run:
	make swag
	go run ./cmd/main.go

swag:
	export PATH=$(go env GOPATH)/bin:$PATH
	swag fmt
	swag init -g cmd/main.go

test:
	go test -v ./tests