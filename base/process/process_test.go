package process

import (
	"runtime"
	"testing"
)

func TestForLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i)
	}
}

func TestWhileLoop(t *testing.T) {
	i := 0
	for i < 5 {
		i++
		t.Log(i)
	}
}

func TestIfCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			t.Log(i)
		}
	}

	if res, err := 1, 0; err == 0 {
		t.Log("success", res)
	} else {
		t.Log("fail", err)
	}
}

func TestSwitchCondition(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("Mac")
	case "linux":
		t.Log("Linux")
	case "windows":
		t.Log("Windows")
	default:
		t.Log(os)
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6, 8, 10:
			t.Log("Even", i)
		case 1, 3, 5, 7, 9:
			t.Log("odd", i)
		default:
			t.Log("default", i)
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even", i)
		case i%2 == 1:
			t.Log("odd", i)
		default:
			t.Log("default", i)
		}
	}
}
