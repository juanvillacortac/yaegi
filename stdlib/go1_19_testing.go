// Code generated by 'yaegi extract testing'. DO NOT EDIT.

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package stdlib

import (
	"reflect"
	"testing"
)

func init() {
	Symbols["testing/testing"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AllocsPerRun":  reflect.ValueOf(testing.AllocsPerRun),
		"Benchmark":     reflect.ValueOf(testing.Benchmark),
		"CoverMode":     reflect.ValueOf(testing.CoverMode),
		"Coverage":      reflect.ValueOf(testing.Coverage),
		"Init":          reflect.ValueOf(testing.Init),
		"Main":          reflect.ValueOf(testing.Main),
		"MainStart":     reflect.ValueOf(testing.MainStart),
		"RegisterCover": reflect.ValueOf(testing.RegisterCover),
		"RunBenchmarks": reflect.ValueOf(testing.RunBenchmarks),
		"RunExamples":   reflect.ValueOf(testing.RunExamples),
		"RunTests":      reflect.ValueOf(testing.RunTests),
		"Short":         reflect.ValueOf(testing.Short),
		"Verbose":       reflect.ValueOf(testing.Verbose),

		// type definitions
		"B":                  reflect.ValueOf((*testing.B)(nil)),
		"BenchmarkResult":    reflect.ValueOf((*testing.BenchmarkResult)(nil)),
		"Cover":              reflect.ValueOf((*testing.Cover)(nil)),
		"CoverBlock":         reflect.ValueOf((*testing.CoverBlock)(nil)),
		"F":                  reflect.ValueOf((*testing.F)(nil)),
		"InternalBenchmark":  reflect.ValueOf((*testing.InternalBenchmark)(nil)),
		"InternalExample":    reflect.ValueOf((*testing.InternalExample)(nil)),
		"InternalFuzzTarget": reflect.ValueOf((*testing.InternalFuzzTarget)(nil)),
		"InternalTest":       reflect.ValueOf((*testing.InternalTest)(nil)),
		"M":                  reflect.ValueOf((*testing.M)(nil)),
		"PB":                 reflect.ValueOf((*testing.PB)(nil)),
		"T":                  reflect.ValueOf((*testing.T)(nil)),
		"TB":                 reflect.ValueOf((*testing.TB)(nil)),

		// interface wrapper definitions
		"_TB": reflect.ValueOf((*_testing_TB)(nil)),
	}
}

// _testing_TB is an interface wrapper for TB type
type _testing_TB struct {
	IValue   interface{}
	WCleanup func(a0 func())
	WError   func(args ...any)
	WErrorf  func(format string, args ...any)
	WFail    func()
	WFailNow func()
	WFailed  func() bool
	WFatal   func(args ...any)
	WFatalf  func(format string, args ...any)
	WHelper  func()
	WLog     func(args ...any)
	WLogf    func(format string, args ...any)
	WName    func() string
	WSetenv  func(key string, value string)
	WSkip    func(args ...any)
	WSkipNow func()
	WSkipf   func(format string, args ...any)
	WSkipped func() bool
	WTempDir func() string
}

func (W _testing_TB) Cleanup(a0 func())                 { W.WCleanup(a0) }
func (W _testing_TB) Error(args ...any)                 { W.WError(args...) }
func (W _testing_TB) Errorf(format string, args ...any) { W.WErrorf(format, args...) }
func (W _testing_TB) Fail()                             { W.WFail() }
func (W _testing_TB) FailNow()                          { W.WFailNow() }
func (W _testing_TB) Failed() bool                      { return W.WFailed() }
func (W _testing_TB) Fatal(args ...any)                 { W.WFatal(args...) }
func (W _testing_TB) Fatalf(format string, args ...any) { W.WFatalf(format, args...) }
func (W _testing_TB) Helper()                           { W.WHelper() }
func (W _testing_TB) Log(args ...any)                   { W.WLog(args...) }
func (W _testing_TB) Logf(format string, args ...any)   { W.WLogf(format, args...) }
func (W _testing_TB) Name() string                      { return W.WName() }
func (W _testing_TB) Setenv(key string, value string)   { W.WSetenv(key, value) }
func (W _testing_TB) Skip(args ...any)                  { W.WSkip(args...) }
func (W _testing_TB) SkipNow()                          { W.WSkipNow() }
func (W _testing_TB) Skipf(format string, args ...any)  { W.WSkipf(format, args...) }
func (W _testing_TB) Skipped() bool                     { return W.WSkipped() }
func (W _testing_TB) TempDir() string                   { return W.WTempDir() }
