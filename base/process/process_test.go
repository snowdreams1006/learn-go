package process

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"testing"
)

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

func TestIfConditionMultiReturnValue(t *testing.T) {
	const filename = "test.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Log(err)
	} else {
		t.Logf("%s\n", content)
	}
}

func TestIfConditionMultiReturnValueShorter(t *testing.T) {
	const filename = "test.txt"
	if content, err := ioutil.ReadFile(filename); err != nil {
		t.Log(err)
	} else {
		t.Logf("%s\n", content)
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

func evalBySwitchOperator(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func TestEvalBySwitchOperator(t *testing.T) {
	// 3
	t.Log(evalBySwitchOperator(1, 2, "+"))
	// -1
	t.Log(evalBySwitchOperator(1, 2, "-"))
	// 2
	t.Log(evalBySwitchOperator(1, 2, "*"))
	// 0
	t.Log(evalBySwitchOperator(1, 2, "/"))
	// unsupported operator:% [recovered]
	//t.Log(evalBySwitchOperator(1, 2, "%"))
}

func gradeBySwitchOperator(score int) string {
	result := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		result = "F"
	case score < 80:
		result = "C"
	case score < 90:
		result = "B"
	case score <= 100:
		result = "A"
	}
	return result
}

func TestGradeBySwitchOperator(t *testing.T) {
	// F F C C B B A A
	t.Log(
		gradeBySwitchOperator(0),
		gradeBySwitchOperator(59),
		gradeBySwitchOperator(60),
		gradeBySwitchOperator(79),
		gradeBySwitchOperator(80),
		gradeBySwitchOperator(89),
		gradeBySwitchOperator(99),
		gradeBySwitchOperator(100),
		//gradeBySwitchOperator(1000),
	)
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

func TestForLoop(t *testing.T) {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	// 1+2+3+...+99+100=5050
	t.Log(sum)
}

func convert2Binary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func TestConvert2Binary(t *testing.T) {
	// 1 100 101 1101
	t.Log(
		convert2Binary(1),
		convert2Binary(4),
		convert2Binary(5),
		convert2Binary(13),
	)
}

func TestWhileLoopByForLoop(t *testing.T) {
	i := 0
	for i < 5 {
		i++
		t.Log(i)
	}
}

func printFile(filename string) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

func TestPrintFile(t *testing.T){
	const filename = "test.txt"
	printFile(filename)
}