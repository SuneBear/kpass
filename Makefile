test:
	go test --race ./app
	go test --race ./app/crypto
	go test --race ./app/dao
	go test --race ./app/dao/user
	go test --race ./app/pkg
	go test --race ./app/api/user
	go test --race ./app/api/entry

build:
	go-bindata -ignore=\\.DS_Store -o ./app/bindata.go -pkg app web/...

reset:
	go-bindata -ignore=\\.* -o ./app/bindata.go -pkg app web/...

cover:
	rm -f *.coverprofile
	go test -coverprofile=app.coverprofile ./app
	gover
	go tool cover -html=gover.coverprofile
	rm -f *.coverprofile

doc:
	godoc -http=:6060

.PHONY: test build cover doc
