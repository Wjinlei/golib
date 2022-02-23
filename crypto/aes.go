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

	"github.com/axgle/mahonia"
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

func AesCFBEncrypt(plaintext []byte, key string, keyLength int, paddingType ...string) (ciphertext []byte, err error) {
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
			plaintext = ZeroPadding(plaintext, aes.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, aes.BlockSize)
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

func AesCFBEncryptToHex(plaintext []byte, key string, keyLength int, paddingType ...string) (string, error) {
	cipherText, err := AesCFBEncrypt(plaintext, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func EncodeStringAesCFBEncryptToHex(plaintext string, key string, keyLength int, coding string, paddingType ...string) (string, error) {
	enCoder := mahonia.NewEncoder(coding)
	if enCoder == nil {
		return "", fmt.Errorf("GetEncoder failed")
	}
	hexText, err := AesCFBEncryptToHex([]byte(enCoder.ConvertString(plaintext)), key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hexText, nil
}

func AesCFBDecrypt(ciphertext []byte, key string, keyLength int, paddingType ...string) (plaintext []byte, err error) {
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
			plaintext = ZeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = ZeroUnPadding(ciphertext)
		}
	} else {
		plaintext = PKCS5UnPadding(ciphertext)
	}
	return plaintext, nil
}

func AesCFBDecryptFromHex(hextext string, key string, keyLength int, paddingType ...string) (string, error) {
	deCodeHex, err := hex.DecodeString(hextext)
	if err != nil {
		return "", err
	}
	plainText, err := AesCFBDecrypt(deCodeHex, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func DecodeStringAesCFBDecryptFromHex(hextext string, key string, keyLength int, coding string, paddingType ...string) (string, error) {
	plainText, err := AesCFBDecryptFromHex(hextext, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	deCoder := mahonia.NewDecoder(coding)
	if deCoder == nil {
		return "", fmt.Errorf("GetDecoder failed")
	}
	return deCoder.ConvertString(plainText), nil
}

func AesCBCEncrypt(plaintext []byte, key string, keyLength int, paddingType ...string) (ciphertext []byte, err error) {
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
			plaintext = ZeroPadding(plaintext, aes.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	}
	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func AesCBCEncryptToHex(plaintext []byte, key string, keyLength int, paddingType ...string) (string, error) {
	cipherText, err := AesCBCEncrypt(plaintext, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func EncodeStringAesCBCEncryptToHex(plaintext string, key string, keyLength int, coding string, paddingType ...string) (string, error) {
	enCoder := mahonia.NewEncoder(coding)
	if enCoder == nil {
		return "", fmt.Errorf("GetEncoder failed")
	}
	hexText, err := AesCBCEncryptToHex([]byte(enCoder.ConvertString(plaintext)), key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hexText, nil
}

func AesCBCDecrypt(ciphertext []byte, key string, keyLength int, paddingType ...string) (plaintext []byte, err error) {
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
			plaintext = ZeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = PKCS5UnPadding(ciphertext)
		}
	} else {
		plaintext = PKCS5UnPadding(ciphertext)
	}
	return plaintext, nil
}

func AesCBCDecryptFromHex(hextext string, key string, keyLength int, paddingType ...string) (string, error) {
	deCodeHex, err := hex.DecodeString(hextext)
	if err != nil {
		return "", err
	}
	plainText, err := AesCBCDecrypt(deCodeHex, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func DecodeStringAesCBCDecryptFromHex(hextext string, key string, keyLength int, coding string, paddingType ...string) (string, error) {
	plainText, err := AesCBCDecryptFromHex(hextext, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	deCoder := mahonia.NewDecoder(coding)
	if deCoder == nil {
		return "", fmt.Errorf("GetDecoder failed")
	}
	return deCoder.ConvertString(plainText), nil
}

func AesECBEncrypt(plaintext []byte, key string, keyLength int, paddingType ...string) (ciphertext []byte, err error) {
	keyByte, err := setAesKey(key, keyLength)
	if err != nil {
		return nil, err
	}
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroPadding":
			plaintext = ZeroPadding(plaintext, aes.BlockSize)
		case "PKCS5Padding":
			plaintext = PKCS5Padding(plaintext, aes.BlockSize)
		}
	} else {
		plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	}
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return nil, err
	}
	ciphertext = make([]byte, len(plaintext))
	NewECBEncrypter(block).CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

func AesECBEncryptToHex(plaintext []byte, key string, keyLength int, paddingType ...string) (string, error) {
	cipherText, err := AesECBEncrypt(plaintext, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cipherText), nil
}

func EncodeStringAesECBEncryptToHex(plaintext string, key string, keyLength int, coding string, paddingType ...string) (string, error) {
	enCoder := mahonia.NewEncoder(coding)
	if enCoder == nil {
		return "", fmt.Errorf("GetEncoder failed")
	}
	hexText, err := AesECBEncryptToHex([]byte(enCoder.ConvertString(plaintext)), key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return hexText, nil
}

func AesECBDecrypt(ciphertext []byte, key string, keyLength int, paddingType ...string) (plaintext []byte, err error) {
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
	NewECBDecrypter(block).CryptBlocks(ciphertext, ciphertext)
	if len(paddingType) > 0 {
		switch paddingType[0] {
		case "ZeroUnPadding":
			plaintext = ZeroUnPadding(ciphertext)
		case "PKCS5UnPadding":
			plaintext = PKCS5UnPadding(ciphertext)
		}
	} else {
		plaintext = PKCS5UnPadding(ciphertext)
	}
	return plaintext, nil
}

func AesECBDecryptFromHex(hextext string, key string, keyLength int, paddingType ...string) (string, error) {
	deCodeHex, err := hex.DecodeString(hextext)
	if err != nil {
		return "", err
	}
	plainText, err := AesECBDecrypt(deCodeHex, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

func DecodeStringAesECBDecryptFromHex(hextext string, key string, keyLength int, coding string, paddingType ...string) (string, error) {
	plainText, err := AesECBDecryptFromHex(hextext, key, keyLength, paddingType...)
	if err != nil {
		return "", err
	}
	deCoder := mahonia.NewDecoder(coding)
	if deCoder == nil {
		return "", fmt.Errorf("GetDecoder failed")
	}
	return deCoder.ConvertString(plainText), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
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

type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
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

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
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
