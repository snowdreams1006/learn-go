package pack

import (
	"fmt"
	"testing"
)

type MyLang struct {
	l *Lang
}

func (ml *MyLang) Print() {
	if ml == nil || ml.l == nil {
		return
	}

	fmt.Println(ml.l.ToString())
}

func TestMyLangPrint(t *testing.T) {
	var l = new(Lang)
	l.SetName("Go")
	l.SetWebsite("https://golang.google.cn/")

	var ml = MyLang{l}

	ml.Print()
}

func TestPack(t *testing.T) {
	var l = new(Lang)
	l.SetName("Go")
	l.SetWebsite("https://golang.google.cn/")

	t.Log(l.ToString())
}

type Lan Lang

func (l *Lan) PrintWebsite() {
	fmt.Println(l.website)
}

func TestLanPrintWebsite(t *testing.T) {
	var la = new(Lan)
	la.name = "GoLang"
	la.website = "https://golang.google.cn/"

	la.PrintWebsite()
}

func TestInitLang(t *testing.T) {
	l := Lang{
		name:    "Go",
		website: "https://golang.google.cn/",
	}

	t.Log(l.ToString())
}
