package main

import (
	"fmt"
	"github.com/snowdreams1006/learn-go/oop/pack"
)

func main() {
	var l = new(pack.Lang)
	l.SetName("Go")
	l.SetWebsite("https://golang.google.cn/")

	fmt.Println(l.ToString())

	l.PrintLangName()
}
