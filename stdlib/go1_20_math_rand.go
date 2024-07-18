// Code generated by 'yaegi extract math/rand'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package stdlib

import (
	"math/rand"
	"reflect"
)

func init() {
	Symbols["math/rand/rand"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ExpFloat64":  reflect.ValueOf(rand.ExpFloat64),
		"Float32":     reflect.ValueOf(rand.Float32),
		"Float64":     reflect.ValueOf(rand.Float64),
		"Int":         reflect.ValueOf(rand.Int),
		"Int31":       reflect.ValueOf(rand.Int31),
		"Int31n":      reflect.ValueOf(rand.Int31n),
		"Int63":       reflect.ValueOf(rand.Int63),
		"Int63n":      reflect.ValueOf(rand.Int63n),
		"Intn":        reflect.ValueOf(rand.Intn),
		"New":         reflect.ValueOf(rand.New),
		"NewSource":   reflect.ValueOf(rand.NewSource),
		"NewZipf":     reflect.ValueOf(rand.NewZipf),
		"NormFloat64": reflect.ValueOf(rand.NormFloat64),
		"Perm":        reflect.ValueOf(rand.Perm),
		"Read":        reflect.ValueOf(rand.Read),
		"Seed":        reflect.ValueOf(rand.Seed),
		"Shuffle":     reflect.ValueOf(rand.Shuffle),
		"Uint32":      reflect.ValueOf(rand.Uint32),
		"Uint64":      reflect.ValueOf(rand.Uint64),

		// type definitions
		"Rand":     reflect.ValueOf((*rand.Rand)(nil)),
		"Source":   reflect.ValueOf((*rand.Source)(nil)),
		"Source64": reflect.ValueOf((*rand.Source64)(nil)),
		"Zipf":     reflect.ValueOf((*rand.Zipf)(nil)),

		// interface wrapper definitions
		"_Source":   reflect.ValueOf((*_math_rand_Source)(nil)),
		"_Source64": reflect.ValueOf((*_math_rand_Source64)(nil)),
	}
}

// _math_rand_Source is an interface wrapper for Source type
type _math_rand_Source struct {
	IValue interface{}
	WInt63 func() int64
	WSeed  func(seed int64)
}

func (W _math_rand_Source) Int63() int64    { return W.WInt63() }
func (W _math_rand_Source) Seed(seed int64) { W.WSeed(seed) }

// _math_rand_Source64 is an interface wrapper for Source64 type
type _math_rand_Source64 struct {
	IValue  interface{}
	WInt63  func() int64
	WSeed   func(seed int64)
	WUint64 func() uint64
}

func (W _math_rand_Source64) Int63() int64    { return W.WInt63() }
func (W _math_rand_Source64) Seed(seed int64) { W.WSeed(seed) }
func (W _math_rand_Source64) Uint64() uint64  { return W.WUint64() }
