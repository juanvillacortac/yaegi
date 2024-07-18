// Code generated by 'yaegi extract image/jpeg'. DO NOT EDIT.

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package stdlib

import (
	"go/constant"
	"go/token"
	"image/jpeg"
	"reflect"
)

func init() {
	Symbols["image/jpeg/jpeg"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Decode":         reflect.ValueOf(jpeg.Decode),
		"DecodeConfig":   reflect.ValueOf(jpeg.DecodeConfig),
		"DefaultQuality": reflect.ValueOf(constant.MakeFromLiteral("75", token.INT, 0)),
		"Encode":         reflect.ValueOf(jpeg.Encode),

		// type definitions
		"FormatError":      reflect.ValueOf((*jpeg.FormatError)(nil)),
		"Options":          reflect.ValueOf((*jpeg.Options)(nil)),
		"Reader":           reflect.ValueOf((*jpeg.Reader)(nil)),
		"UnsupportedError": reflect.ValueOf((*jpeg.UnsupportedError)(nil)),

		// interface wrapper definitions
		"_Reader": reflect.ValueOf((*_image_jpeg_Reader)(nil)),
	}
}

// _image_jpeg_Reader is an interface wrapper for Reader type
type _image_jpeg_Reader struct {
	IValue    interface{}
	WRead     func(p []byte) (n int, err error)
	WReadByte func() (byte, error)
}

func (W _image_jpeg_Reader) Read(p []byte) (n int, err error) { return W.WRead(p) }
func (W _image_jpeg_Reader) ReadByte() (byte, error)          { return W.WReadByte() }
