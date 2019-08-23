package pack

import (
	"testing"
)

func TestPack(t *testing.T){
	var l = new(Lang)
	l.SetName("Go")
	l.SetWebsite("https://golang.google.cn/")

	t.Log(l.ToString())
}