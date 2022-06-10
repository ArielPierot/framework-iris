LINUX_AMD64 = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
PROJECT_NAME = desafio-golang

install-deps:
	go mod tidy
	go mod download

build-binaries:
	$(LINUX_AMD64) go build -o $(PROJECT_NAME) main.go

install-linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2

lint:
	golangci-lint run ./...

test:
	go test -covermode=count -coverprofile=count.out ./...

build:
	$(LINUX_AMD64) go build -o iris-framework main.go
