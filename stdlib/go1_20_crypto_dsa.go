// Code generated by 'yaegi extract crypto/dsa'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package stdlib

import (
	"crypto/dsa"
	"reflect"
)

func init() {
	Symbols["crypto/dsa/dsa"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrInvalidPublicKey": reflect.ValueOf(&dsa.ErrInvalidPublicKey).Elem(),
		"GenerateKey":         reflect.ValueOf(dsa.GenerateKey),
		"GenerateParameters":  reflect.ValueOf(dsa.GenerateParameters),
		"L1024N160":           reflect.ValueOf(dsa.L1024N160),
		"L2048N224":           reflect.ValueOf(dsa.L2048N224),
		"L2048N256":           reflect.ValueOf(dsa.L2048N256),
		"L3072N256":           reflect.ValueOf(dsa.L3072N256),
		"Sign":                reflect.ValueOf(dsa.Sign),
		"Verify":              reflect.ValueOf(dsa.Verify),

		// type definitions
		"ParameterSizes": reflect.ValueOf((*dsa.ParameterSizes)(nil)),
		"Parameters":     reflect.ValueOf((*dsa.Parameters)(nil)),
		"PrivateKey":     reflect.ValueOf((*dsa.PrivateKey)(nil)),
		"PublicKey":      reflect.ValueOf((*dsa.PublicKey)(nil)),
	}
}
