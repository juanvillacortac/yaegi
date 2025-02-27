// Code generated by 'yaegi extract net/http/httptest'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package stdlib

import (
	"go/constant"
	"go/token"
	"net/http/httptest"
	"reflect"
)

func init() {
	Symbols["net/http/httptest/httptest"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"DefaultRemoteAddr":  reflect.ValueOf(constant.MakeFromLiteral("\"1.2.3.4\"", token.STRING, 0)),
		"NewRecorder":        reflect.ValueOf(httptest.NewRecorder),
		"NewRequest":         reflect.ValueOf(httptest.NewRequest),
		"NewServer":          reflect.ValueOf(httptest.NewServer),
		"NewTLSServer":       reflect.ValueOf(httptest.NewTLSServer),
		"NewUnstartedServer": reflect.ValueOf(httptest.NewUnstartedServer),

		// type definitions
		"ResponseRecorder": reflect.ValueOf((*httptest.ResponseRecorder)(nil)),
		"Server":           reflect.ValueOf((*httptest.Server)(nil)),
	}
}
