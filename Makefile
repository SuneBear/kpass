test: assets
	go test --race ./pkg
	go test --race ./pkg/api
	go test --race ./pkg/auth
	go test --race ./pkg/dao
	go test --race ./pkg/logger
	go test --race ./pkg/schema
	go test --race ./pkg/service
	go test --race ./pkg/util

cover: assets
	rm -f *.coverprofile
	go test -coverprofile=pkg.coverprofile ./pkg
	go test -coverprofile=api.coverprofile ./pkg/api
	go test -coverprofile=auth.coverprofile ./pkg/auth
	go test -coverprofile=dao.coverprofile ./pkg/dao
	go test -coverprofile=logger.coverprofile ./pkg/logger
	go test -coverprofile=schema.coverprofile ./pkg/schema
	go test -coverprofile=service.coverprofile ./pkg/service
	go test -coverprofile=util.coverprofile ./pkg/util
	gover
	go tool cover -html=gover.coverprofile
	rm -f *.coverprofile
	make clean

GO=$(shell which go)

assets:
	go-bindata -ignore=\\.DS_Store -o ./pkg/bindata.go -pkg pkg web/...
clean:
	go-bindata -ignore=\\.* -o ./pkg/bindata.go -pkg pkg web/...

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo -o dist/kpass_linux ./cmd/kpass
build-osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -a -installsuffix cgo -o dist/kpass ./cmd/kpass
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build -a -installsuffix cgo -o dist/kpass.exe ./cmd/kpass
build: assets build-osx build-linux build-win clean

.PHONY: assets test build cover clean
