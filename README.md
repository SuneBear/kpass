# KPass
====
KPass is a web application to manage password safe.

[![Build Status](http://img.shields.io/travis/seccom/kpass.svg?style=flat-square)](https://travis-ci.org/seccom/kpass)
[![Coverage Status](http://img.shields.io/coveralls/seccom/kpass.svg?style=flat-square)](https://coveralls.io/r/seccom/kpass)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/seccom/kpass/master/LICENSE)

## Build

```sh
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

A demo user `{id:"demo", pass:"demo"}` will be created for a new database.