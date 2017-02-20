test: assets
	go test --race ./src
	go test --race ./src/api
	go test --race ./src/auth
	go test --race ./src/dao
	go test --race ./src/logger
	go test --race ./src/schema
	go test --race ./src/service
	go test --race ./src/util

cover: assets
	rm -f *.coverprofile
	go test -coverprofile=src.coverprofile ./src
	go test -coverprofile=api.coverprofile ./src/api
	go test -coverprofile=auth.coverprofile ./src/auth
	go test -coverprofile=dao.coverprofile ./src/dao
	go test -coverprofile=logger.coverprofile ./src/logger
	go test -coverprofile=schema.coverprofile ./src/schema
	go test -coverprofile=service.coverprofile ./src/service
	go test -coverprofile=util.coverprofile ./src/util
	gover
	go tool cover -html=gover.coverprofile
	rm -f *.coverprofile
	make clean

GO=$(shell which go)

assets:
	go-bindata -ignore=\\.DS_Store -o ./src/bindata.go -pkg src web/dist/...
clean:
	go-bindata -ignore=\\.* -o ./src/bindata.go -pkg src web/dist/...

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo -o dist/kpass_linux ./cmd/kpass
build-osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -a -installsuffix cgo -o dist/kpass ./cmd/kpass
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build -a -installsuffix cgo -o dist/kpass.exe ./cmd/kpass
build: assets build-osx build-linux build-win clean

.PHONY: assets test build cover clean
