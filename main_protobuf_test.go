package main

import (
	"testing"

	"github.com/Jimeux/chacha/user"
	"github.com/golang/protobuf/proto"
)

func BenchmarkRegularProtobufEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out, _ := proto.Marshal(regularUser)
		_, _ = encrypter.Encrypt(out)
	}
}

func BenchmarkRegularProtobufDecrypt(b *testing.B) {
	out, _ := proto.Marshal(regularUser)
	ev, _ := encrypter.Encrypt(out)

	for i := 0; i < b.N; i++ {
		cp := make([]byte, len(ev))
		copy(cp, ev)

		raw, _ := encrypter.Decrypt(cp)
		u := new(user.User)
		_ = proto.Unmarshal(raw, u)
	}
}

func BenchmarkBigProtobufEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		out, _ := proto.Marshal(bigUser)
		_, _ = encrypter.Encrypt(out)
	}
}

func BenchmarkBigProtobufDecrypt(b *testing.B) {
	out, _ := proto.Marshal(bigUser)
	ev, _ := encrypter.Encrypt(out)

	for i := 0; i < b.N; i++ {
		cp := make([]byte, len(ev))
		copy(cp, ev)

		raw, _ := encrypter.Decrypt(cp)
		u := new(user.User)
		_ = proto.Unmarshal(raw, u)
	}
}
