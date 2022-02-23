package ssl

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// Decode 解码SSL加密字符串
func Decode(encode string) (*pem.Block, error) {
	pemBlock, _ := pem.Decode([]byte(encode))
	if pemBlock == nil {
		return nil, errors.New("解码失败,请确认pem或key是否填写正确")
	}
	return pemBlock, nil
}

// ParseCertificate 解析pemBlock为x509.Certificate
func ParseCertificate(pemBlock *pem.Block) (*x509.Certificate, error) {
	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

// ParseKey 解析keyBlock为key
func ParseKey(keyBlock *pem.Block) (interface{}, error) {
	key, err := x509.ParsePKCS8PrivateKey(keyBlock.Bytes)
	if err != nil {
		key, err = x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
		if err != nil {
			return nil, err
		}
	}
	return key, nil
}

// IsMatch 判断pem和key是否匹配
func IsMatch(cert *x509.Certificate, key interface{}) bool {
	if cert.PublicKey.(*rsa.PublicKey).N.String() != key.(*rsa.PrivateKey).PublicKey.N.String() {
		return false
	}
	return true
}
