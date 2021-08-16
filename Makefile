code-check: mod imports fmt vet
vet:
	go vet ./...
fmt:
	gofmt -d -s .
imports:
	goimports -w .
build:
	sh docker-build.sh
mod:
	go mod tidy
	go mod verify
	go mod download


install-oapi-codegen:
	go get -u github.com/deepmap/oapi-codegen/cmd/oapi-codegen

oapi-codegen:
	oapi-codegen -generate "types" -package gen petstore-expanded.yaml > ./internal/http/gen/model.go
	oapi-codegen -generate "server,spec" -package gen petstore-expanded.yaml > ./internal/http/gen/server.go
