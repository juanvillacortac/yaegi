// Code generated by 'yaegi extract net/http/fcgi'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package stdlib

import (
	"net/http/fcgi"
	"reflect"
)

func init() {
	Symbols["net/http/fcgi/fcgi"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrConnClosed":     reflect.ValueOf(&fcgi.ErrConnClosed).Elem(),
		"ErrRequestAborted": reflect.ValueOf(&fcgi.ErrRequestAborted).Elem(),
		"ProcessEnv":        reflect.ValueOf(fcgi.ProcessEnv),
		"Serve":             reflect.ValueOf(fcgi.Serve),
	}
}
