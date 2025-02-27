// Code generated by 'yaegi extract crypto/cipher'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package stdlib

import (
	"crypto/cipher"
	"reflect"
)

func init() {
	Symbols["crypto/cipher/cipher"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewCBCDecrypter":     reflect.ValueOf(cipher.NewCBCDecrypter),
		"NewCBCEncrypter":     reflect.ValueOf(cipher.NewCBCEncrypter),
		"NewCFBDecrypter":     reflect.ValueOf(cipher.NewCFBDecrypter),
		"NewCFBEncrypter":     reflect.ValueOf(cipher.NewCFBEncrypter),
		"NewCTR":              reflect.ValueOf(cipher.NewCTR),
		"NewGCM":              reflect.ValueOf(cipher.NewGCM),
		"NewGCMWithNonceSize": reflect.ValueOf(cipher.NewGCMWithNonceSize),
		"NewGCMWithTagSize":   reflect.ValueOf(cipher.NewGCMWithTagSize),
		"NewOFB":              reflect.ValueOf(cipher.NewOFB),

		// type definitions
		"AEAD":         reflect.ValueOf((*cipher.AEAD)(nil)),
		"Block":        reflect.ValueOf((*cipher.Block)(nil)),
		"BlockMode":    reflect.ValueOf((*cipher.BlockMode)(nil)),
		"Stream":       reflect.ValueOf((*cipher.Stream)(nil)),
		"StreamReader": reflect.ValueOf((*cipher.StreamReader)(nil)),
		"StreamWriter": reflect.ValueOf((*cipher.StreamWriter)(nil)),

		// interface wrapper definitions
		"_AEAD":      reflect.ValueOf((*_crypto_cipher_AEAD)(nil)),
		"_Block":     reflect.ValueOf((*_crypto_cipher_Block)(nil)),
		"_BlockMode": reflect.ValueOf((*_crypto_cipher_BlockMode)(nil)),
		"_Stream":    reflect.ValueOf((*_crypto_cipher_Stream)(nil)),
	}
}

// _crypto_cipher_AEAD is an interface wrapper for AEAD type
type _crypto_cipher_AEAD struct {
	IValue     interface{}
	WNonceSize func() int
	WOpen      func(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error)
	WOverhead  func() int
	WSeal      func(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte
}

func (W _crypto_cipher_AEAD) NonceSize() int { return W.WNonceSize() }
func (W _crypto_cipher_AEAD) Open(dst []byte, nonce []byte, ciphertext []byte, additionalData []byte) ([]byte, error) {
	return W.WOpen(dst, nonce, ciphertext, additionalData)
}
func (W _crypto_cipher_AEAD) Overhead() int { return W.WOverhead() }
func (W _crypto_cipher_AEAD) Seal(dst []byte, nonce []byte, plaintext []byte, additionalData []byte) []byte {
	return W.WSeal(dst, nonce, plaintext, additionalData)
}

// _crypto_cipher_Block is an interface wrapper for Block type
type _crypto_cipher_Block struct {
	IValue     interface{}
	WBlockSize func() int
	WDecrypt   func(dst []byte, src []byte)
	WEncrypt   func(dst []byte, src []byte)
}

func (W _crypto_cipher_Block) BlockSize() int                 { return W.WBlockSize() }
func (W _crypto_cipher_Block) Decrypt(dst []byte, src []byte) { W.WDecrypt(dst, src) }
func (W _crypto_cipher_Block) Encrypt(dst []byte, src []byte) { W.WEncrypt(dst, src) }

// _crypto_cipher_BlockMode is an interface wrapper for BlockMode type
type _crypto_cipher_BlockMode struct {
	IValue       interface{}
	WBlockSize   func() int
	WCryptBlocks func(dst []byte, src []byte)
}

func (W _crypto_cipher_BlockMode) BlockSize() int                     { return W.WBlockSize() }
func (W _crypto_cipher_BlockMode) CryptBlocks(dst []byte, src []byte) { W.WCryptBlocks(dst, src) }

// _crypto_cipher_Stream is an interface wrapper for Stream type
type _crypto_cipher_Stream struct {
	IValue        interface{}
	WXORKeyStream func(dst []byte, src []byte)
}

func (W _crypto_cipher_Stream) XORKeyStream(dst []byte, src []byte) { W.WXORKeyStream(dst, src) }
