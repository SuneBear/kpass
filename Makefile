test:
	go test --race ./server
	go test --race ./server/api
	go test --race ./server/crypto
	go test --race ./server/dao

build:
	go-bindata -ignore=\\.DS_Store -o ./server/bindata.go -pkg app web/...

reset:
	go-bindata -ignore=\\.* -o ./server/bindata.go -pkg app web/...

cover:
	rm -f *.coverprofile
	go test -coverprofile=server.coverprofile ./server
	gover
	go tool cover -html=gover.coverprofile
	rm -f *.coverprofile

doc:
	godoc -http=:6060

.PHONY: test build cover doc
