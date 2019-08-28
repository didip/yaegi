// Code generated by 'goexports net/mail'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlib

import (
	"net/mail"
	"reflect"
)

func init() {
	Symbols["net/mail"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrHeaderNotPresent": reflect.ValueOf(&mail.ErrHeaderNotPresent).Elem(),
		"ParseAddress":        reflect.ValueOf(mail.ParseAddress),
		"ParseAddressList":    reflect.ValueOf(mail.ParseAddressList),
		"ParseDate":           reflect.ValueOf(mail.ParseDate),
		"ReadMessage":         reflect.ValueOf(mail.ReadMessage),

		// type definitions
		"Address":       reflect.ValueOf((*mail.Address)(nil)),
		"AddressParser": reflect.ValueOf((*mail.AddressParser)(nil)),
		"Header":        reflect.ValueOf((*mail.Header)(nil)),
		"Message":       reflect.ValueOf((*mail.Message)(nil)),
	}
}