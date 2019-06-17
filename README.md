# chacha

#### Sample code for using [package chacha20poly1305](https://godoc.org/golang.org/x/crypto/chacha20poly1305) with benchmarks.

### `encrypt` package

`Encrypter` is a simple utility type using the package that provides the following features:

- Protects against double encryption / decrypting unencrypted values
- Generates 24-byte nonce values per encryption
- Supports a simple mechanism for switching cipher keys

### Benchmarks

The benchmarks compare serializing a struct with JSON and Protobuf before encrypting, and do the reverse for decryption. 

#### User struct with short values

Operation | Time
---|---
Proto marshal + encrypt | 2.3-2.6ms
JSON marshal + encrypt | 3-3.4ms
Proto unmarshal + decrypt | 1-1.3ms
JSON unmarshal + decrypt | 4.3-4.7ms 💩

#### User struct with long values

Operation | Time
---|---
Proto marshal + encrypt | 4.6-5.4ms
JSON marshal + encrypt | 6.4-7.1ms
Proto unmarshal + decrypt | 3.3-3.6ms
JSON unmarshal + decrypt | 14-15ms 💩💩💩💩💩
