# KPass

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
const globalHMACFn = (data) => HMAC(SHA256, dbSalt)(data)
const globalAESKeyFn = (a, b) => base64Encode(globalHMACFn(globalHMACFn(a) + globalHMACFn(b)))
const globalPBKDF2Fn = (data, iv) => PBKDF2(dbSalt, 12480, 64, HMAC(SHA512, iv))(data)
const globalEncryptFn = (key, data) => {
  let cipherData = AES-CTR(globalHMACFn(key), IV(16), data)
  let sum = HMAC(SHA1, cipherData)(data)
  return cipherData + sum
}
const globalDecryptFn = reverse(globalEncryptFn)
```

### User password

It is used to verify user.

```js
let UserPass = SHA256("someUserPassword")
let data = globalHmac(UserID) + UserPass
let iv = IV(8)
let data = globalPBKDF2Fn(data, iv)
let UserCheckPass = base64Encode(data + iv)
// Save UserCheckPass to user Model
```

### User AES Key

It is used to encrypt and decrypt user's data.

```js
let UserAESKey = globalAESKeyFn(UserPass, UserCheckPass)
```

### Team password

It is used to generate TeamKey.

```js
let TeamPass = SHA256(RandPass(20))
let data = globalHmac(TeamID) + TeamPass
let iv = IV(8)
let data = globalPBKDF2Fn(data, iv)
let TeamCheckPass = base64Encode(data + iv)
// Save TeamCheckPass to team Model
```

### Team AES Key

It is used to encrypt and decrypt secret messages and files in team' entris.

```js
let TeamAESKey = globalAESKeyFn(TeamPass, TeamCheckPass)
```

### Team password for member

All team members should able to get TeamAESKey to encrypt and decrypt.

**When user login and create a team:**

```js
let CipherTeamPass = globalEncryptFn(UserAESKey, TeamPass)
// Save CipherTeamPass to database with TeamID and UserID
```

**When user login and read or write team's data:**

```js
let UserAESKey = globalAESKeyFn(UserPass, UserCheckPass)
let TeamPass = globalDecryptFn(UserAESKey, CipherTeamPass)
let TeamAESKey = globalAESKeyFn(TeamPass, TeamCheckPass)
let cipherData = globalEncryptFn(TeamAESKey, data)
let data = globalDecryptFn(TeamAESKey, cipherData)
```

**When user A login and invite another user B to the team:**

```js
let UserAESKey_A = globalAESKeyFn(UserPass_A, UserCheckPass_A)
let TeamPass = globalDecryptFn(UserAESKey_A, CipherTeamPass)
let AESKey = globalAESKeyFn(UserCheckPass_A, UserCheckPass_B)
let InviteCode = globalEncryptFn(AESKey, TeamPass)
// user A send InviteCode to user B, user B logined
let UserAESKey_B = globalAESKeyFn(UserPass_B, UserCheckPass_B)
let AESKey = globalAESKeyFn(UserCheckPass_A, UserCheckPass_B)
let TeamPass = globalDecryptFn(AESKey, InviteCode)
// Check TeamPass with TeamCheckPass
let CipherTeamPass = globalEncryptFn(UserAESKey_B, TeamPass)
// Save CipherTeamPass to database with TeamID and UserID_B
```