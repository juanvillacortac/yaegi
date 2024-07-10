// Code generated by 'yaegi extract testing/quick'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package stdlib

import (
	"math/rand"
	"reflect"
	"testing/quick"
)

func init() {
	Symbols["testing/quick/quick"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Check":      reflect.ValueOf(quick.Check),
		"CheckEqual": reflect.ValueOf(quick.CheckEqual),
		"Value":      reflect.ValueOf(quick.Value),

		// type definitions
		"CheckEqualError": reflect.ValueOf((*quick.CheckEqualError)(nil)),
		"CheckError":      reflect.ValueOf((*quick.CheckError)(nil)),
		"Config":          reflect.ValueOf((*quick.Config)(nil)),
		"Generator":       reflect.ValueOf((*quick.Generator)(nil)),
		"SetupError":      reflect.ValueOf((*quick.SetupError)(nil)),

		// interface wrapper definitions
		"_Generator": reflect.ValueOf((*_testing_quick_Generator)(nil)),
	}
}

// _testing_quick_Generator is an interface wrapper for Generator type
type _testing_quick_Generator struct {
	IValue    interface{}
	WGenerate func(rand *rand.Rand, size int) reflect.Value
}

func (W _testing_quick_Generator) Generate(rand *rand.Rand, size int) reflect.Value {
	return W.WGenerate(rand, size)
}
