// Code generated by 'yaegi extract go/doc/comment'. DO NOT EDIT.

//go:build go1.21
// +build go1.21

package stdlib

import (
	"go/doc/comment"
	"reflect"
)

func init() {
	Symbols["go/doc/comment/comment"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"DefaultLookupPackage": reflect.ValueOf(comment.DefaultLookupPackage),

		// type definitions
		"Block":     reflect.ValueOf((*comment.Block)(nil)),
		"Code":      reflect.ValueOf((*comment.Code)(nil)),
		"Doc":       reflect.ValueOf((*comment.Doc)(nil)),
		"DocLink":   reflect.ValueOf((*comment.DocLink)(nil)),
		"Heading":   reflect.ValueOf((*comment.Heading)(nil)),
		"Italic":    reflect.ValueOf((*comment.Italic)(nil)),
		"Link":      reflect.ValueOf((*comment.Link)(nil)),
		"LinkDef":   reflect.ValueOf((*comment.LinkDef)(nil)),
		"List":      reflect.ValueOf((*comment.List)(nil)),
		"ListItem":  reflect.ValueOf((*comment.ListItem)(nil)),
		"Paragraph": reflect.ValueOf((*comment.Paragraph)(nil)),
		"Parser":    reflect.ValueOf((*comment.Parser)(nil)),
		"Plain":     reflect.ValueOf((*comment.Plain)(nil)),
		"Printer":   reflect.ValueOf((*comment.Printer)(nil)),
		"Text":      reflect.ValueOf((*comment.Text)(nil)),

		// interface wrapper definitions
		"_Block": reflect.ValueOf((*_go_doc_comment_Block)(nil)),
		"_Text":  reflect.ValueOf((*_go_doc_comment_Text)(nil)),
	}
}

// _go_doc_comment_Block is an interface wrapper for Block type
type _go_doc_comment_Block struct {
	IValue interface{}
}

// _go_doc_comment_Text is an interface wrapper for Text type
type _go_doc_comment_Text struct {
	IValue interface{}
}
