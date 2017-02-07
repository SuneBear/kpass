KPass
====
KPass is a web application to manage password safe.

[![Build Status](http://img.shields.io/travis/seccom/kpass.svg?style=flat-square)](https://travis-ci.org/seccom/kpass)
[![Coverage Status](http://img.shields.io/coveralls/seccom/kpass.svg?style=flat-square)](https://coveralls.io/r/seccom/kpass)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/seccom/kpass/master/LICENSE)

## Build

```sh
go get -t github.com/seccom/kpass
go get -u github.com/jteeuwen/go-bindata/...
cd $GOPATH/src/github.com/seccom/kpass
make build
```
It will build three executable files for OSX, windows and linux version in "./dist" directory.

### Run in OSX
```sh
cd ./dist
./kpass --help
./kpass
```
It will run with default options, create a `kpass.kdb` file and open a browser.

### Development

Start a development mode with memory database:
```sh
make assets
go run cmd/kpass/kpass.go -dev
```
It creates some demo data. You can find the encrypted secret in the `kpass.kdb`.

It will serve `./web` as static server too.