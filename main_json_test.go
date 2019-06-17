package main

import (
	"encoding/json"
	"testing"

	"github.com/Jimeux/chacha/encrypt"
	"github.com/Jimeux/chacha/user"
)

var (
	regularUser = &user.User{
		FirstName:          "太郎",
		SecondName:         "山田",
		PhoneticFirstName:  "たろう",
		PhoneticSecondName: "やまだ",
		Email:              "taro.yamada@gmail.com",
	}
	bigUser = &user.User{
		FirstName:          "漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字漢字",
		SecondName:         "多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数多数",
		PhoneticFirstName:  "かんじかんじかんじかんじかんじかんじかんじかんじかんじかんじかんじかんじかんじかんじかんじかんじかんじ",
		PhoneticSecondName: "たすうたすうたすうたすうたすうたすうたすうたすうたすうたすうたすうたすうたすうたすうたすうたすうたすう",
		Email:              "manymanymanymanymanymanymanymanymany.kanjikanjikanjikanjikanjikanjikanjikanji@gmail.com",
	}
	encrypter, _ = encrypt.NewEncrypter(encrypt.KeyVersion(1), encrypt.KeyMap{
		encrypt.KeyVersion(1): encrypt.Key("itWouldBeBadIfSomebodyFoundThis!"),
	})
)

func BenchmarkRegularJSONEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byt, _ := json.Marshal(regularUser)
		_, _ = encrypter.Encrypt(byt)
	}
}

func BenchmarkRegularJSONDecrypt(b *testing.B) {
	byt, _ := json.Marshal(regularUser)
	ev, _ := encrypter.Encrypt(byt)

	for i := 0; i < b.N; i++ {
		// アロケーション＋コピーもベンチマークに影響を与えてしまう
		cp := make([]byte, len(ev))
		copy(cp, ev)

		raw, _ := encrypter.Decrypt(cp)
		u := new(user.User)
		_ = json.Unmarshal(raw, u)
	}
}

func BenchmarkBigJSONEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byt, _ := json.Marshal(bigUser)
		_, _ = encrypter.Encrypt(byt)
	}
}

func BenchmarkBigJSONDecrypt(b *testing.B) {
	byt, _ := json.Marshal(bigUser)
	ev, _ := encrypter.Encrypt(byt)

	for i := 0; i < b.N; i++ {
		// アロケーション＋コピーもベンチマークに影響を与えてしまう
		cp := make([]byte, len(ev))
		copy(cp, ev)

		raw, _ := encrypter.Decrypt(cp)
		u := new(user.User)
		_ = json.Unmarshal(raw, u)
	}
}
