// Code generated by 'yaegi extract crypto/subtle'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package stdlib

import (
	"crypto/subtle"
	"reflect"
)

func init() {
	Symbols["crypto/subtle/subtle"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ConstantTimeByteEq":   reflect.ValueOf(subtle.ConstantTimeByteEq),
		"ConstantTimeCompare":  reflect.ValueOf(subtle.ConstantTimeCompare),
		"ConstantTimeCopy":     reflect.ValueOf(subtle.ConstantTimeCopy),
		"ConstantTimeEq":       reflect.ValueOf(subtle.ConstantTimeEq),
		"ConstantTimeLessOrEq": reflect.ValueOf(subtle.ConstantTimeLessOrEq),
		"ConstantTimeSelect":   reflect.ValueOf(subtle.ConstantTimeSelect),
		"XORBytes":             reflect.ValueOf(subtle.XORBytes),
	}
}
