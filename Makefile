test: build-assets
	go test --race ./app
	go test --race ./app/api/entry
	go test --race ./app/api/user
	go test --race ./app/crypto
	go test --race ./app/dao
	go test --race ./app/dao/entry
	go test --race ./app/dao/secret
	go test --race ./app/dao/share
	go test --race ./app/dao/team
	go test --race ./app/dao/user
	go test --race ./app/pkg

cover:
	rm -f *.coverprofile
	go test -coverprofile=app.coverprofile ./app
	gover
	go tool cover -html=gover.coverprofile
	rm -f *.coverprofile

doc:
	godoc -http=:6060

GO=$(shell which go)

build-assets:
	go-bindata -ignore=\\.DS_Store -o ./app/bindata.go -pkg app web/...
clean:
	go-bindata -ignore=\\.* -o ./app/bindata.go -pkg app web/...

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -installsuffix cgo -o dist/kpass_linux .
build-osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -a -installsuffix cgo -o dist/kpass .
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build -a -installsuffix cgo -o dist/kpass.exe .
build: build-assets build-osx build-linux build-win clean

.PHONY: test build cover doc
