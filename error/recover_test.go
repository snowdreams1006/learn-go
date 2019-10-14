package main

import (
	"errors"
	"testing"
)

func TestPanic(t *testing.T) {
	// 「雪之梦技术驿站」: This is an error
	panic(errors.New("「雪之梦技术驿站」: This is an error"))
}