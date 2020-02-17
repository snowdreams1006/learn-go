package strings

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string

	// s = ,len(s) =  0
	t.Logf("s = %v,len(s) =  %d", s, len(s))

	s = "hello"

	// s = hello,len(s) =  5
	t.Logf("s = %v,len(s) =  %d", s, len(s))

	s = "\xE4\xB8\xA5"

	// s = 严,len(s) =  3
	t.Logf("s = %v,len(s) =  %d", s, len(s))

	s = "\xE444\xB888\xA555"

	// s = �44�88�55,len(s) =  9
	t.Logf("s = %v,len(s) =  %d", s, len(s))
}

func TestStringImmutable(t *testing.T) {
	var s string

	// s = ,len(s) =  0
	t.Logf("s = %v,len(s) =  %d", s, len(s))

	s = "hello"

	// s = hello,len(s) =  5
	t.Logf("s = %v,len(s) =  %d", s, len(s))

	// cannot assign to s[0]
	//s[0] = 1
}

func TestStringCodePoint(t *testing.T) {
	var s string

	// s = ,len(s) =  0
	t.Logf("s = %v,len(s) =  %d", s, len(s))

	s = "hello"

	// s = hello,len(s) =  5
	t.Logf("s = %v,len(s) =  %d", s, len(s))

	s = "中"

	// s = 中,len(s) =  3
	t.Logf("s = %v,len(s) =  %d", s, len(s))

	//s = e4b8ad,len(s) =  3
	t.Logf("s = %x,len(s) =  %d", s, len(s))

	ch := []rune(s)

	// s = [4e2d],len(s) =  1
	t.Logf("ch = %x,len(ch) =  %d", ch, len(ch))
}

func TestString2Rune(t *testing.T) {
	s := "中华人民共和国"

	for _, c := range s {
		t.Logf("%[1]c %[1]d %[1]x", c)
	}
}

func TestStrings(t *testing.T) {
	s := "A,B,C"

	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	newStr := strings.Join(parts, "-")

	t.Log(newStr)
}

func TestStrconv(t *testing.T) {

	s := strconv.Itoa(10)

	// string 10
	t.Logf("%[1]T %[1]v", s)

	// int 10
	if i, err := strconv.Atoi("10"); err == nil {
		t.Logf("%[1]T %[1]v", i)
	}
}
