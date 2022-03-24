package crypto

// +---------------------------------------------------------------------
// | Description: AES加解密
// +---------------------------------------------------------------------
// | Copyright (c) 2004-2020 护卫神(http://hws.com) All rights reserved.
// +---------------------------------------------------------------------
// | Author: Wjinlei <1976883731@qq.com>
// +---------------------------------------------------------------------
//
//                  ___====-_  _-====___
//             _--^^^#####/      \#####^^^--_
//          _-^##########/ (    ) \##########^-_
//         -############/  |\^^/|  \############-
//       _/############/   (@::@)   \############\_
//     /#############((     \  /     ))#############\
//     -###############\    (oo)    /###############-
//    -#################\  / VV \  /#################-
//   -###################\/      \/###################-
// _#/|##########/\######(   /\   )######/\##########|\#_
// |/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
// '  |/  V  V      V  \#\| |  | |/#/  V      V  V  \|  '
//    '   '  '      '   / | |  | | \   '      '  '   '
//                     (  | |  | |  )
//                    __\ | |  | | /__
//                   (vvv(VVV)(VVV)vvv)
//
//                  神龙护体
//                代码无bug!

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/Wjinlei/golib/encoding"
)

func setAesKey(key string, length int) ([]byte, error) {
	b := make([]byte, length)
	c := []byte(key)
	copy(b, c)
	//128 192 256位的其中一个长度,分别对应16 24 32字节长度
	if len(b) == 16 || len(b) == 24 || len(b) == 32 {
		return b, nil
	}
	return nil, fmt.Errorf("key size is not 16 or 24 or 32, but %d", len(b))
}

func aesCFBEncrypt(plaintext []byte, key string, keyLength int, paddingType ...string) (ciphertext []byte, err error) {
	keyByte, err := setAesKey(key, keyLength)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = zeroPadding(plaintext, aes.BlockSize)
		case "PKCS5Padding":
			plaintext = pkcs5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = pkcs5Padding(plaintext, aes.BlockSize)
	}

	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:],
		plaintext)
	return ciphertext, nil

}

func GbkAesCFBEncrypt(plaintext string, key string, keyLength int, paddingType ...string) (string, error) {
	plainTextGbk, err := encoding.Utf8ToGbk([]byte(plaintext))
	if err != nil {
		return "", err
	}
	cipherText, err := aesCFBEncrypt(plainTextGbk, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func Utf8AesCFBEncrypt(plaintext string, key string, keyLength int, paddingType ...string) (string, error) {
	cipherText, err := aesCFBEncrypt([]byte(plaintext), key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func aesCFBDecrypt(ciphertext []byte, key string, keyLength int, paddingType ...string) (plaintext []byte, err error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	keyByte, err := setAesKey(key, keyLength)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(ciphertext, ciphertext)
	if int(ciphertext[len(ciphertext)-1]) > len(ciphertext) {
		return nil, errors.New("aes decrypt failed")
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroUnPadding":
			plaintext = zeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = zeroUnPadding(ciphertext)
		}
	} else {
		plaintext = pkcs5UnPadding(ciphertext)
	}
	return plaintext, nil
}

func GbkAesCFBDecrypt(hexText string, key string, keyLength int, paddingType ...string) (string, error) {
	deCodeHex, err := hex.DecodeString(hexText)
	if err != nil {
		return "", err
	}
	plainText, err := aesCFBDecrypt(deCodeHex, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	plainTextUtf8, err := encoding.GbkToUtf8(plainText)
	if err != nil {
		return "", err
	}
	return string(plainTextUtf8), nil
}

func Utf8AesCFBDecrypt(hexText string, key string, keyLength int, paddingType ...string) (string, error) {
	deCodeHex, err := hex.DecodeString(hexText)
	if err != nil {
		return "", err
	}
	plainText, err := aesCFBDecrypt(deCodeHex, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func aesCBCEncrypt(plaintext []byte, key string, keyLength int, paddingType ...string) (ciphertext []byte, err error) {
	keyByte, err := setAesKey(key, keyLength)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = zeroPadding(plaintext, aes.BlockSize)
		case "PKCS5Padding":
			plaintext = pkcs5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = pkcs5Padding(plaintext, aes.BlockSize)
	}
	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func GbkAesCBCEncrypt(plaintext string, key string, keyLength int, paddingType ...string) (string, error) {
	plainTextGbk, err := encoding.Utf8ToGbk([]byte(plaintext))
	if err != nil {
		return "", err
	}
	cipherText, err := aesCBCEncrypt(plainTextGbk, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func Utf8AesCBCEncrypt(plaintext string, key string, keyLength int, paddingType ...string) (string, error) {
	cipherText, err := aesCBCEncrypt([]byte(plaintext), key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func aesCBCDecrypt(ciphertext []byte, key string, keyLength int, paddingType ...string) (plaintext []byte, err error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	keyByte, err := setAesKey(key, keyLength)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(ciphertext, ciphertext)
	if int(ciphertext[len(ciphertext)-1]) > len(ciphertext) {
		return nil, errors.New("aes decrypt failed")
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroUnPadding":
			plaintext = zeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = pkcs5UnPadding(ciphertext)
		}
	} else {
		plaintext = pkcs5UnPadding(ciphertext)
	}
	return plaintext, nil
}

func GbkAesCBCDecrypt(hexText string, key string, keyLength int, paddingType ...string) (string, error) {
	deCodeHex, err := hex.DecodeString(hexText)
	if err != nil {
		return "", err
	}
	plainText, err := aesCBCDecrypt(deCodeHex, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	plainTextUtf8, err := encoding.GbkToUtf8(plainText)
	if err != nil {
		return "", err
	}
	return string(plainTextUtf8), nil
}

func Utf8AesCBCDecrypt(hexText string, key string, keyLength int, paddingType ...string) (string, error) {
	deCodeHex, err := hex.DecodeString(hexText)
	if err != nil {
		return "", err
	}
	plainText, err := aesCBCDecrypt(deCodeHex, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func aesECBEncrypt(plaintext []byte, key string, keyLength int, paddingType ...string) (ciphertext []byte, err error) {
	keyByte, err := setAesKey(key, keyLength)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = zeroPadding(plaintext, aes.BlockSize)
		case "PKCS5Padding":
			plaintext = pkcs5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = pkcs5Padding(plaintext, aes.BlockSize)
	}
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	ciphertext = make([]byte, len(plaintext))
	newECBEncryptor(block).CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

func GbkAesECBEncrypt(plaintext string, key string, keyLength int, paddingType ...string) (string, error) {
	plainTextGbk, err := encoding.Utf8ToGbk([]byte(plaintext))
	if err != nil {
		return "", err
	}
	cipherText, err := aesECBEncrypt(plainTextGbk, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func Utf8AesECBEncrypt(plaintext string, key string, keyLength int, paddingType ...string) (string, error) {
	cipherText, err := aesECBEncrypt([]byte(plaintext), key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func aesECBDecrypt(ciphertext []byte, key string, keyLength int, paddingType ...string) (plaintext []byte, err error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	keyByte, err := setAesKey(key, keyLength)
	if err != nil {
		return nil, err
	}
	// ECB mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	newECBDecryptor(block).CryptBlocks(ciphertext, ciphertext)
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroUnPadding":
			plaintext = zeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = pkcs5UnPadding(ciphertext)
		}
	} else {
		plaintext = pkcs5UnPadding(ciphertext)
	}
	return plaintext, nil
}

func GbkAesECBDecrypt(hexText string, key string, keyLength int, paddingType ...string) (string, error) {
	deCodeHex, err := hex.DecodeString(hexText)
	if err != nil {
		return "", err
	}
	plainText, err := aesECBDecrypt(deCodeHex, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	plainTextUtf8, err := encoding.GbkToUtf8(plainText)
	if err != nil {
		return "", err
	}
	return string(plainTextUtf8), nil
}

func Utf8AesECBDecrypt(hexText string, key string, keyLength int, paddingType ...string) (string, error) {
	deCodeHex, err := hex.DecodeString(hexText)
	if err != nil {
		return "", err
	}
	plainText, err := aesECBDecrypt(deCodeHex, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncryptor ecb

func newECBEncryptor(b cipher.Block) cipher.BlockMode {
	return (*ecbEncryptor)(newECB(b))
}

func (x *ecbEncryptor) BlockSize() int { return x.blockSize }

func (x *ecbEncryptor) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecryptor ecb

func newECBDecryptor(b cipher.Block) cipher.BlockMode {
	return (*ecbDecryptor)(newECB(b))
}

func (x *ecbDecryptor) BlockSize() int { return x.blockSize }

func (x *ecbDecryptor) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
