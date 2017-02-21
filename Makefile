test:
	APP_ENV=test go test --race ./src
	APP_ENV=test go test --race ./src/api
	APP_ENV=test go test --race ./src/auth
	APP_ENV=test go test --race ./src/dao
	APP_ENV=test go test --race ./src/logger
	APP_ENV=test go test --race ./src/schema
	APP_ENV=test go test --race ./src/service
	APP_ENV=test go test --race ./src/util

cover:
	rm -f *.coverprofile
	APP_ENV=test go test -coverprofile=src.coverprofile ./src
	APP_ENV=test go test -coverprofile=api.coverprofile ./src/api
	APP_ENV=test go test -coverprofile=auth.coverprofile ./src/auth
	APP_ENV=test go test -coverprofile=dao.coverprofile ./src/dao
	APP_ENV=test go test -coverprofile=logger.coverprofile ./src/logger
	APP_ENV=test go test -coverprofile=schema.coverprofile ./src/schema
	APP_ENV=test go test -coverprofile=service.coverprofile ./src/service
	APP_ENV=test go test -coverprofile=util.coverprofile ./src/util
	gover
	go tool cover -html=gover.coverprofile
	rm -f *.coverprofile

GO=$(shell which go)

assets:
	go-bindata -ignore=\\.DS_Store -o ./src/bindata.go -pkg src -prefix web/dist/ web/dist/...
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
