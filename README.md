# KPass

KPass is a web application to manage password safe.

[![Build Status](http://img.shields.io/travis/seccom/kpass.svg?style=flat-square)](https://travis-ci.org/seccom/kpass)
[![Coverage Status](http://img.shields.io/coveralls/seccom/kpass.svg?style=flat-square)](https://coveralls.io/r/seccom/kpass)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/seccom/kpass/master/LICENSE)

## Feature

1. Support multi-users
1. Support multi-teams
1. Support HTTPS and HTTP/2
1. Support secret files(TODO)
1. Share secret to other user(TODO)

## Build

```sh
go get -u github.com/seccom/kpass
go get -u github.com/jteeuwen/go-bindata/...
cd $GOPATH/src/github.com/seccom/kpass
cd web
yarn install
cd -
make build
```

It will build three executable files for OSX, windows and linux version in "./dist" directory.

### Run in OSX

```sh
./dist/kpass --help
./dist/kpass
```

It will run with default options, create a `kpass.kdb` file and open a browser.

### Development

Start a development mode with memory database:

```sh
make dev
```

It creates some demo data. You can find the encrypted secret in the `kpass.kdb`.

It will serve `./web` as static server too.

### Swagger Document

```sh
go install github.com/teambition/swaggo
go install github.com/teambition/gear/example/staticgo
```

```sh
make swagger
open http://petstore.swagger.io/?url=http://127.0.0.1:3000/swagger.json
```

## Security Design

```js
globalHMACFn = (a, b) => HMAC(SHA256, a)(b)
globalAESKeyFn = (a, b) => base64Encode(globalHMACFn(a + b))
globalPBKDF2Fn = (data, iv) => PBKDF2(dbSalt, 12480, 64, HMAC(SHA512, iv))(data)
globalEncryptFn = (key, data) => {
  cipherData = AESCTR(globalHMACFn(key), IV(16), data)
  sum = HMAC(SHA1, cipherData)(data)
  return cipherData + sum
}
globalDecryptFn = reverse(globalEncryptFn)
```

### User password

It is used to verify user.

```js
UserPass = SHA256("someUserPassword")
data = globalHmac(UserID) + UserPass
iv = IV(8)
data = globalPBKDF2Fn(data, iv)
UserCheckPass = base64Encode(data + iv)
// Save UserCheckPass to user Model
```

### User AES Key

It is used to encrypt and decrypt user's data.

```js
UserAESKey = globalAESKeyFn(UserPass, UserCheckPass)
```

### Team password

It is used to generate TeamKey.

```js
TeamPass = SHA256(RandPass(20))
data = globalHmac(TeamID) + TeamPass
iv = IV(8)
data = globalPBKDF2Fn(data, iv)
TeamCheckPass = base64Encode(data + iv)
// Save TeamCheckPass to team Model
```

### Team AES Key

It is used to encrypt and decrypt secret messages and files in team' entris.

```js
TeamAESKey = globalAESKeyFn(TeamPass, TeamCheckPass)
```

### Team password for member

All team members should able to get TeamAESKey to encrypt and decrypt.

**When user login and create a team:**

```js
CipherTeamPass = globalEncryptFn(UserAESKey, TeamPass)
// Save CipherTeamPass to database with TeamID and UserID
```

**When user login and read or write team's data:**

```js
UserAESKey = globalAESKeyFn(UserPass, UserCheckPass)
TeamPass = globalDecryptFn(UserAESKey, CipherTeamPass)
TeamAESKey = globalAESKeyFn(TeamPass, TeamCheckPass)
cipherData = globalEncryptFn(TeamAESKey, data)
data = globalDecryptFn(TeamAESKey, cipherData)
```

**When user A login and invite another user B to the team:**

```js
UserAESKey_A = globalAESKeyFn(UserPass_A, UserCheckPass_A)
TeamPass = globalDecryptFn(UserAESKey_A, CipherTeamPass)
AESKey = globalAESKeyFn(UserCheckPass_A, UserCheckPass_B)
InviteCode = globalEncryptFn(AESKey, TeamPass)
// user A send InviteCode to user B, user B logined
UserAESKey_B = globalAESKeyFn(UserPass_B, UserCheckPass_B)
AESKey = globalAESKeyFn(UserCheckPass_A, UserCheckPass_B)
TeamPass = globalDecryptFn(AESKey, InviteCode)
// Check TeamPass with TeamCheckPass
CipherTeamPass = globalEncryptFn(UserAESKey_B, TeamPass)
// Save CipherTeamPass to database with TeamID and UserID_B
```