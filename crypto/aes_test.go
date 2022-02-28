package crypto

import (
	"fmt"
	"testing"
)

func TestGbkAesECBEncrypt(t *testing.T) {
	encrypt, err := GbkAesECBEncrypt("Hello 中国!", "123", 16)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(encrypt)
}

func TestGbkAesECBDecrypt(t *testing.T) {
	decrypt, err := GbkAesECBDecrypt("24ece90186db56d04172f46b934f6e33", "123", 16)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(decrypt)
}

func TestUtf8AesECBEncrypt(t *testing.T) {
	encrypt, err := Utf8AesECBEncrypt("Hello 中国!", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(encrypt)
}

func TestUtf8AesECBDecrypt(t *testing.T) {
	decrypt, err := Utf8AesECBDecrypt("c073f512e975de0ac675e4c2c3607020", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decrypt)
}

func TestGbkAesCBCEncrypt(t *testing.T) {
	encrypt, err := GbkAesCBCEncrypt("Hello 中国!", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(encrypt)
}

func TestGbkAesCBCDecrypt(t *testing.T) {
	decrypt, err := GbkAesCBCDecrypt("54c8bac8259002f11efea988d6161103ab8f93fd57fbe4d5c4ccd045513111f2", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decrypt)
}

func TestUtf8AesCBCEncrypt(t *testing.T) {
	encrypt, err := Utf8AesCBCEncrypt("Hello 中国!", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(encrypt)
}

func TestUtf8AesCBCDecrypt(t *testing.T) {
	decrypt, err := Utf8AesCBCDecrypt("7b7f54840856016a62ee098c5c23899c7ffbf63f3b33d63bd7b97cbcf3da5479", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decrypt)
}

func TestGbkAesCFBEncrypt(t *testing.T) {
	encrypt, err := GbkAesCFBEncrypt("Hello 中国!", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(encrypt)
}

func TestGbkAesCFBDecrypt(t *testing.T) {
	decrypt, err := GbkAesCFBDecrypt("1ea9fc345e83cfe6ac8ea61ddfb757411b12c2cda9d0ede6442601da355ec633", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decrypt)
}

func TestUtf8AesCFBEncrypt(t *testing.T) {
	encrypt, err := Utf8AesCFBEncrypt("Hello 中国!", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(encrypt)
}

func TestUtf8AesCFBDecrypt(t *testing.T) {
	decrypt, err := Utf8AesCFBDecrypt("a748e27c2911e95b03f7bb87e55356016926c6fb660f13d0ebd2c8090aaf63e3", "123", 16)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decrypt)
}
