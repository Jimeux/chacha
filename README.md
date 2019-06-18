# chacha

#### Sample code for using [package chacha20poly1305](https://godoc.org/golang.org/x/crypto/chacha20poly1305) with benchmarks.

### `encrypt` package

`Encrypter` is a simple utility type using the package that provides the following features:

- Protects against double encryption / decrypting unencrypted values
- Generates 24-byte nonce values per encryption
- Supports a simple mechanism for switching cipher keys

### Benchmarks

The benchmarks compare encrypting struct values serialized with JSON and Protobuf. 

#### User struct with short values

Function | Executions | Time
---|---|---
BenchmarkRegularJSONEncrypt-8 |	1000000	| 2070 ns/op
BenchmarkRegularJSONDecrypt-8       |	  300000	 |   3370 ns/op
BenchmarkRegularProtobufEncrypt-8   |	 1000000	 |   1418 ns/op
BenchmarkRegularProtobufDecrypt-8   |	 2000000	  |   722 ns/op

#### User struct with long values

Function | Executions | Time
---|---|---
BenchmarkBigJSONEncrypt-8           |	  300000	 |   4345 ns/op
BenchmarkBigJSONDecrypt-8         |  	  200000 |   10114 ns/op
BenchmarkBigProtobufEncrypt-8       |	  500000	 |   3094 ns/op
BenchmarkBigProtobufDecrypt-8       |	  500000	 |   2350 ns/op
