// Code generated by 'goexports net/http/fcgi'. DO NOT EDIT.

// +build go1.11,!go1.12

package stdlib

import (
	"net/http/fcgi"
	"reflect"
)

func init() {
	Symbols["net/http/fcgi"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ErrConnClosed":     reflect.ValueOf(&fcgi.ErrConnClosed).Elem(),
		"ErrRequestAborted": reflect.ValueOf(&fcgi.ErrRequestAborted).Elem(),
		"ProcessEnv":        reflect.ValueOf(fcgi.ProcessEnv),
		"Serve":             reflect.ValueOf(fcgi.Serve),
	}
}