// Code generated by 'yaegi extract container/heap'. DO NOT EDIT.

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package stdlib

import (
	"container/heap"
	"reflect"
)

func init() {
	Symbols["container/heap/heap"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Fix":    reflect.ValueOf(heap.Fix),
		"Init":   reflect.ValueOf(heap.Init),
		"Pop":    reflect.ValueOf(heap.Pop),
		"Push":   reflect.ValueOf(heap.Push),
		"Remove": reflect.ValueOf(heap.Remove),

		// type definitions
		"Interface": reflect.ValueOf((*heap.Interface)(nil)),

		// interface wrapper definitions
		"_Interface": reflect.ValueOf((*_container_heap_Interface)(nil)),
	}
}

// _container_heap_Interface is an interface wrapper for Interface type
type _container_heap_Interface struct {
	IValue interface{}
	WLen   func() int
	WLess  func(i int, j int) bool
	WPop   func() any
	WPush  func(x any)
	WSwap  func(i int, j int)
}

func (W _container_heap_Interface) Len() int               { return W.WLen() }
func (W _container_heap_Interface) Less(i int, j int) bool { return W.WLess(i, j) }
func (W _container_heap_Interface) Pop() any               { return W.WPop() }
func (W _container_heap_Interface) Push(x any)             { W.WPush(x) }
func (W _container_heap_Interface) Swap(i int, j int)      { W.WSwap(i, j) }
