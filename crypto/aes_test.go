package crypto

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// TestEncodeStringAesECBEncryptToHex 将字符串转换为指定编码,再进行AES ECB模式加密,最后编码为十六进制字符串
func TestEncodeStringAesECBEncryptToHex(t *testing.T) {
	encrypt, err := EncodeStringAesECBEncryptToHex("Hello 中国!", "123", 16, "gbk")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(encrypt)
}

// TestDecodeStringAesECBDecryptFromHex 从十六进制字符串进行AES ECB模式解密,然后解码为指定编码的字符串
func TestDecodeStringAesECBDecryptFromHex(t *testing.T) {
	decrypt, err := DecodeStringAesECBDecryptFromHex("24ece90186db56d04172f46b934f6e33", "123", 16, "gbk")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(decrypt)
}

// TestAesECBEncryptToHex 将数据进行AES ECB模式加密,然后编码为十六进制字符串(该方法不支持指定字符串编码)
func TestAesECBEncryptToHex(t *testing.T) {
	encrypt, err := AesECBEncryptToHex([]byte("Hello 中国!"), "123", 16)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(encrypt)
}

// TestAesECBDecryptFromHex 将十六进制字符串进行AES ECB模式解密(该方法不支持指定编码)
func TestAesECBDecryptFromHex(t *testing.T) {
	decrypt, err := AesECBDecryptFromHex("c073f512e975de0ac675e4c2c3607020", "123", 16)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(decrypt)
}

// TestAesECBEncrypt AES ECB模式加密[]byte类型数据,通常用于加密文件
func TestAesECBEncrypt(t *testing.T) {
	fileByte, err := ioutil.ReadFile("testfile.txt")
	if err != nil {
		t.Error(err)
	}
	data, err := AesECBEncrypt(fileByte, "123", 16)
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("testfile_encrypt", data, 0644)
	if err != nil {
		t.Error(err)
	}
}

// TestAesEcbDecrypt AES ECB模式解密[]byte类型的数据,通常用于解密加密文件
func TestAesEcbDecrypt(t *testing.T) {
	fileByte, err := ioutil.ReadFile("testfile_encrypt")
	if err != nil {
		t.Error(err)
	}
	data, err := AesECBDecrypt(fileByte, "123", 16)
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("testfile.txt", data, 0644)
	if err != nil {
		t.Error(err)
	}
}
