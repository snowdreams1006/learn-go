package main

import (
	"errors"
	"testing"
)

func TestPanicWithRecover(t *testing.T) {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			t.Logf("Error occurred : %s", err.Error())
		}else{
			panic(r)
		}
	}()

	// 「雪之梦技术驿站」: This is an error
	panic(errors.New("「雪之梦技术驿站」: This is an error"))
}

func TestPanicWithPanic(t *testing.T) {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			t.Logf("Error occurred : %s", err.Error())
		}else{
			panic(r)
		}
	}()

	// 「雪之梦技术驿站」: This is a custom statement
	panic("「雪之梦技术驿站」: This is a custom statement")
}