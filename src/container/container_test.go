package container

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestArray(t *testing.T) {
	var arr1 [3]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	// [0 0 0] [1 2 3 4 5] [2 4 6 8 10]
	t.Log(arr1, arr2, arr3)

	var grid [3][4]int

	// [[0 0 0 0] [0 0 0 0] [0 0 0 0]]
	t.Log(grid)
}

func TestArrayTraverse(t *testing.T) {
	arr := [...]int{2, 4, 6, 8, 10}

	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for i := range arr {
		t.Log(arr[i])
	}

	for i, v := range arr {
		t.Log(i, v)
	}

	for _, v := range arr {
		t.Log(v)
	}
}

func TestFindMaxInArray(t *testing.T) {
	arr := [...]int{2, 4, 6, 8, 10}

	maxVal, maxIdx := -1, -1
	for i, v := range arr {
		if v > maxVal {
			maxVal, maxIdx = v, i
		}
	}

	t.Log(maxVal, maxIdx)
}

func TestSumForArray(t *testing.T) {
	arr := [...]int{2, 4, 6, 8, 10}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	t.Log(sum)
}

func printArray(arr [5]int) {
	arr[0] = 666
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func TestPrintArray(t *testing.T) {
	var arr1 [3]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	// [0 0 0] [1 2 3 4 5] [2 4 6 8 10]
	t.Log(arr1, arr2, arr3)

	// cannot use arr1 (type [3]int) as type [5]int in argument to printArray
	//printArray(arr1)

	fmt.Println("printArray(arr2)")
	printArray(arr2)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	// [1 2 3 4 5] [2 4 6 8 10]
	t.Log(arr2, arr3)
}

func printArrayByPointer(arr *[5]int) {
	arr[0] = 666
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func TestPrintArrayByPointer(t *testing.T) {
	var arr1 [3]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	// [0 0 0] [1 2 3 4 5] [2 4 6 8 10]
	t.Log(arr1, arr2, arr3)

	fmt.Println("printArrayByPointer(arr2)")
	printArrayByPointer(&arr2)

	fmt.Println("printArrayByPointer(arr3)")
	printArrayByPointer(&arr3)

	// [666 2 3 4 5] [666 4 6 8 10]
	t.Log(arr2, arr3)
}

func TestSliceFromArray(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr =  [0 1 2 3 4 5 6 7 8 9]
	t.Log("arr = ", arr)

	// arr[2:6] =  [2 3 4 5]
	t.Log("arr[2:6] = ", arr[2:6])
	// arr[:6] =  [0 1 2 3 4 5]
	t.Log("arr[:6] = ", arr[:6])
	// arr[2:] =  [2 3 4 5 6 7 8 9]
	t.Log("arr[2:] = ", arr[2:])
	// arr[:] =  [0 1 2 3 4 5 6 7 8 9]
	t.Log("arr[:] = ", arr[:])
}

func TestSliceInit(t *testing.T) {
	var s1 [5]int
	// [0 0 0 0 0]
	t.Log(s1)

	var s2 []int
	// []
	t.Log(s2, len(s2))
}

func TestSliceInitValue(t *testing.T) {
	var s1 = [5]int{1, 3, 5, 7, 9}
	// [1 3 5 7 9]
	t.Log(s1)

	var s2 = []int{1, 3, 5, 7, 9}
	// [1 3 5 7 9]
	t.Log(s2)
}

func TestSliceCompare(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// arr =  [0 1 2 3 4 5 6 7 8 9]
	t.Log("arr = ", arr)

	s1 := arr[2:6]
	// s1 =  [2 3 4 5]
	t.Log("s1 = ", s1)

	s2 := arr[2:6]
	// s1 =  [2 3 4 5]
	t.Log("s2 = ", s2)

	//// invalid operation: s1 == s2 (slice can only be compared to nil)
	//t.Log(s1 == s2)
}

func updateSlice(s []int) {
	s[0] = 666
}

func TestUpdateSlice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// arr =  [0 1 2 3 4 5 6 7 8 9]
	t.Log("arr = ", arr)

	s1 := arr[2:6]
	// s1 =  [2 3 4 5]
	t.Log("s1 = ", s1)

	s2 := arr[:6]
	// s2 =  [0 1 2 3 4 5]
	t.Log("s2 = ", s2)

	updateSlice(s1)
	// s1 =  [666 3 4 5]
	t.Log("s1 = ", s1)
	// arr =  [0 1 666 3 4 5 6 7 8 9]
	t.Log("arr = ", arr)

	updateSlice(s2)
	// s2 =  [666 1 666 3 4 5]
	t.Log("s2 = ", s2)
	// arr =  [666 1 666 3 4 5 6 7 8 9]
	t.Log("arr = ", arr)
}

func TestReSlice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := arr[2:6]
	// s1 =  [2 3 4 5]
	t.Log("s1 = ", s1)

	s1 = s1[2:]
	// s1 =  [4 5]
	t.Log("s1 = ", s1)

	s1 = s1[:6]
	// s1 =  [4 5 6 7 8 9]
	t.Log("s1 = ", s1)
}

func TestSliceOutOfBound(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6]
	// s1 = [2 3 4 5], len(s1) = 4, cap(s1) = 6
	t.Logf("s1 = %v, len(s1) = %d, cap(s1) = %d", s1, len(s1), cap(s1))

	s2 := s1[3:5]
	// s2 = [5 6], len(s2) = 2, cap(s2) = 3
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))

	t.Log(s1[3])
	// index out of range
	//t.Log(s1[4])

	t.Log(s1[3:6])
	// slice bounds out of range
	//t.Log(s1[3:7])
}

func TestSliceDetail(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// arr = [0 1 2 3 4 5 6 7], len(arr) = 8, cap(arr) = 8
	t.Logf("arr = %v, len(arr) = %d, cap(arr) = %d", arr, len(arr), cap(arr))

	s1 := arr[2:6]
	// s1 = [2 3 4 5], len(s1) = 4, cap(s1) = 6
	t.Logf("s1 = %v, len(s1) = %d, cap(s1) = %d", s1, len(s1), cap(s1))

	s2 := s1[3:5]
	// s2 = [5 6], len(s2) = 2, cap(s2) = 3
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))

	// slice bounds out of range
	//t.Log(s1[3:7])
}

func TestSliceAppend(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// arr = [0 1 2 3 4 5 6 7], len(arr) = 8, cap(arr) = 8
	t.Logf("arr = %v, len(arr) = %d, cap(arr) = %d", arr, len(arr), cap(arr))

	s1 := arr[2:6]
	// s1 = [2 3 4 5], len(s1) = 4, cap(s1) = 6
	t.Logf("s1 = %v, len(s1) = %d, cap(s1) = %d", s1, len(s1), cap(s1))

	s2 := s1[3:5]
	// s2 = [5 6], len(s2) = 2, cap(s2) = 3
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	// s3 = [5 6 10], len(s3) = 3, cap(s3) = 3
	t.Logf("s3 = %v, len(s3) = %d, cap(s3) = %d", s3, len(s3), cap(s3))

	s4 := append(s3, 11)
	// s4 = [5 6 10 11], len(s4) = 4, cap(s4) = 6
	t.Logf("s4 = %v, len(s4) = %d, cap(s4) = %d", s4, len(s4), cap(s4))

	s5 := append(s4, 12)
	// s5 = [5 6 10 11 12], len(s5) = 5, cap(s5) = 6
	t.Logf("s5 = %v, len(s5) = %d, cap(s5) = %d", s5, len(s5), cap(s5))

	// arr = [0 1 2 3 4 5 6 10], len(arr) = 8, cap(arr) = 8
	t.Logf("arr = %v, len(arr) = %d, cap(arr) = %d", arr, len(arr), cap(arr))
}

func TestNewSlice(t *testing.T) {
	var s []int
	// [] true
	t.Log(s, s == nil)

	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	// [0 1 2 3 4 5 6 7 8 9]
	t.Log(s)
}

func printSlice(s []int) {
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}

func TestSliceAutoLonger(t *testing.T) {
	var s []int
	// []
	t.Log(s)

	for i := 0; i < 10; i++ {
		s = append(s, i)

		printSlice(s)
	}

	// [0 1 2 3 ...,98,99]
	t.Log(s)

	for i := 0; i < 10; i++ {
		s = s[1:]

		printSlice(s)
	}

	// [0 1 2 3 ...,98,99]
	t.Log(s)
}

func TestMakeSlice(t *testing.T) {
	s1 := make([]int, 10)

	// s1 = [0 0 0 0 0 0 0 0 0 0], len(s1) = 10, cap(s1) = 10
	t.Logf("s1 = %v, len(s1) = %d, cap(s1) = %d", s1, len(s1), cap(s1))

	s2 := make([]int, 10, 32)

	// s2 = [0 0 0 0 0 0 0 0 0 0], len(s2) = 10, cap(s2) = 32
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))
}

func TestNewSliceInitialValue(t *testing.T) {
	s := []int{1, 3, 5, 7, 9}

	// s = [1 3 5 7 9], len(s) = 5, cap(s) = 5
	t.Logf("s = %v, len(s) = %d, cap(s) = %d", s, len(s), cap(s))
}

func TestNewSliceInitialLength(t *testing.T) {
	s := make([]int, 10)

	// s = [0 0 0 0 0 0 0 0 0 0], len(s) = 10, cap(s) = 10
	t.Logf("s = %v, len(s) = %d, cap(s) = %d", s, len(s), cap(s))
}

func TestNewSliceInitialLengthAndCapacity(t *testing.T) {
	s := make([]int, 10, 32)

	// s = [0 0 0 0 0 0 0 0 0 0], len(s) = 10, cap(s) = 32
	t.Logf("s = %v, len(s) = %d, cap(s) = %d", s, len(s), cap(s))
}

func TestCopySlice(t *testing.T) {
	var s1 = []int{1, 3, 5, 7, 9}
	var s2 = make([]int, 10, 32)

	copy(s2, s1)

	// s2 = [1 3 5 7 9 0 0 0 0 0], len(s2) = 10, cap(s2) = 32
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))

	var s3 []int

	copy(s3, s1)

	// s3 = [], len(s3) = 0, cap(s3) = 0
	t.Logf("s3 = %v, len(s3) = %d, cap(s3) = %d", s3, len(s3), cap(s3))
}

func TestDeleteSlice(t *testing.T) {
	s1 := []int{1, 3, 5, 7, 9}
	s2 := make([]int, 10, 32)

	copy(s2, s1)

	// s2 = [1 3 5 7 9 0 0 0 0 0], len(s2) = 10, cap(s2) = 32
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))

	s2 = append(s2[:3], s2[4:]...)

	// s2 = [1 3 5 9 0 0 0 0 0], len(s2) = 9, cap(s2) = 32
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))
}

func TestPopSlice(t *testing.T) {
	s1 := []int{1, 3, 5, 7, 9}
	s2 := make([]int, 10, 32)

	copy(s2, s1)

	// s2 = [1 3 5 7 9 0 0 0 0 0], len(s2) = 10, cap(s2) = 32
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))

	front := s2[0]
	s2 = s2[1:]

	// front = 1
	t.Logf("front = %v", front)
	// s2 = [3 5 7 9 0 0 0 0 0], len(s2) = 9, cap(s2) = 31
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	// tail = 0
	t.Logf("tail = %v", tail)
	// s2 = [3 5 7 9 0 0 0 0], len(s2) = 8, cap(s2) = 31
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))
}

func TestMap(t *testing.T) {
	m := map[string]string{
		"name": "snowdreams1006",
		"site": "https://snowdreams1006.github.io",
	}

	// map[name:snowdreams1006 site:https://snowdreams1006.github.io]
	t.Log(m)
}

func TestMapByMake(t *testing.T) {
	// empty map
	m1 := make(map[string]int)

	// map[] false
	t.Log(m1, m1 == nil)

	// nil
	var m2 map[string]int

	// map[] true
	t.Log(m2, m2 == nil)

	// [0 0 0 0 0] false
	a1 := make([]int, 5)
	t.Log(a1, a1 == nil)

	// [] true
	var a2 []int
	t.Log(a2, a2 == nil)
}

func TestMapTraverse(t *testing.T) {
	m := map[string]string{
		"name": "snowdreams1006",
		"site": "https://snowdreams1006.github.io",
	}

	// map[name:snowdreams1006 site:https://snowdreams1006.github.io]
	t.Log(m)

	for k, v := range m {
		t.Log(k, v)
	}

	t.Log()

	for k := range m {
		t.Log(k)
	}

	t.Log()

	for _, v := range m {
		t.Log(v)
	}
}

func TestMapGetItem(t *testing.T) {
	m := map[string]string{
		"name": "snowdreams1006",
		"site": "https://snowdreams1006.github.io",
	}

	// map[name:snowdreams1006 site:https://snowdreams1006.github.io]
	t.Log(m)

	name := m["name"]

	// snowdreams1006
	t.Log(name)

	author := m["author"]

	// zero value is empty
	t.Log(author)

	m2 := map[string]int{
		"id":    1,
		"score": 100,
	}
	// map[id:1 score:100]
	t.Log(m2)
	// zero value is 0
	t.Log(m2["sco"])

	sco, ok := m2["sco"]

	// 0 false
	t.Log(sco, ok)

	score, ok := m2["score"]

	// 100 true
	t.Log(score, ok)

	if id, ok := m2["id"]; ok {
		t.Log(id)
	} else {
		t.Log("key does not exist ")
	}
}

func TestMapDeleteItem(t *testing.T) {
	m := map[string]string{
		"name": "snowdreams1006",
		"site": "https://snowdreams1006.github.io",
	}

	// map[name:snowdreams1006 site:https://snowdreams1006.github.io]
	t.Log(m)

	delete(m, "name")

	// map[site:https://snowdreams1006.github.io]
	t.Log(m)

	delete(m, "id")

	// map[site:https://snowdreams1006.github.io]
	t.Log(m)
}

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	start, maxLength := 0, 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func TestLengthOfLongestSubstring(t *testing.T) {
	// 3 1 3 0 1 7 11 6
	t.Log(
		lengthOfLongestSubstring("abcabcbb"),
		lengthOfLongestSubstring("bbbbb"),
		lengthOfLongestSubstring("pwwkew"),
		lengthOfLongestSubstring(""),
		lengthOfLongestSubstring("a"),
		lengthOfLongestSubstring("abcdefg"),
		lengthOfLongestSubstring("雪之梦技术驿站"),
		lengthOfLongestSubstring("一零零六"),
	)
}

func TestString2Rune(t *testing.T) {
	s := "hello,雪之梦技术驿站"

	//  hello,雪之梦技术驿站 27
	t.Log(s, len(s))

	// 68 65 6C 6C 6F 2C E9 9B AA E4 B9 8B E6 A2 A6 E6 8A 80 E6 9C AF E9 A9 BF E7 AB 99
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	// ch is rune,utf-8 解码再转 unicode 编码
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	// 13
	t.Log(utf8.RuneCountInString(s))

	bytes := []byte(s)

	// h e l l o , 雪 之 梦 技 术 驿 站
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]

		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// (0 h) (1 e) (2 l) (3 l) (4 o) (5 ,) (6 雪) (7 之) (8 梦) (9 技) (10 术) (11 驿) (12 站)
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}

func lengthOfLongestSubstringInternationalVersion(s string) int {
	lastOccurred := make(map[rune]int)
	start, maxLength := 0, 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func TestLengthOfLongestSubstringInternationalVersion(t *testing.T) {
	// 3 1 3 0 1 7 7 2 8
	t.Log(
		lengthOfLongestSubstringInternationalVersion("abcabcbb"),
		lengthOfLongestSubstringInternationalVersion("bbbbb"),
		lengthOfLongestSubstringInternationalVersion("pwwkew"),
		lengthOfLongestSubstringInternationalVersion(""),
		lengthOfLongestSubstringInternationalVersion("a"),
		lengthOfLongestSubstringInternationalVersion("abcdefg"),
		lengthOfLongestSubstringInternationalVersion("雪之梦技术驿站"),
		lengthOfLongestSubstringInternationalVersion("一零零六"),
		lengthOfLongestSubstringInternationalVersion("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"),
	)
}

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}

	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}

	// 1 4 27
	t.Log(m[1](1), m[2](2), m[3](3))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}

	mySet[1] = true

	n := 3

	if mySet[n] {
		t.Log("update", mySet[n])
	} else {
		t.Log("add", mySet[n])
	}

	delete(mySet, 1)
}

func TestArrayAndSlice(t *testing.T) {
	var arr1 [3]int
	var arr2 []int

	// [0 0 0] []
	t.Log(arr1, arr2)
}

func TestArrayAndSliceByUpdate(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// arr =  [0 1 2 3 4 5 6 7 8 9]
	t.Log("arr = ", arr)

	s := arr[2:6]

	// before update s = [2 3 4 5], arr = [0 1 2 3 4 5 6 7 8 9]
	t.Logf("before update s = %v, arr = %v", s, arr)

	s[0] = 666

	// after update s = [666 3 4 5], arr = [0 1 666 3 4 5 6 7 8 9]
	t.Logf("after update s = %v, arr = %v", s, arr)
}

func TestArrayAndSliceIncreasing(t *testing.T) {
	var s []int

	fmt.Println("add new item to slice")

	for i := 0; i < 10; i++ {
		s = append(s, i)

		fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
	}

	fmt.Println("remove item from slice")

	for i := 0; i < 10; i++ {
		s = s[1:]

		fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
	}
}

func TestArrayAndSliceBound(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s1 := arr[5:8]

	// s1[0] = 5, s1[2] = 7
	t.Logf("s1[0] = %d, s1[%d] = %d", s1[0], len(s1)-1, s1[len(s1)-1])
	// s1 = [5 6 7], len(s1) = 3, cap(s1) = 5
	t.Logf("s1 = %v, len(s1) = %d, cap(s1) = %d", s1, len(s1), cap(s1))

	s2 := s1[3:5]

	// s2[0] = 8, s2[1] = 9
	t.Logf("s2[0] = %d, s2[%d] = %d", s2[0], len(s2)-1, s2[len(s2)-1])
	// s2 = [8 9], len(s2) = 2, cap(s2) = 2
	t.Logf("s2 = %v, len(s2) = %d, cap(s2) = %d", s2, len(s2), cap(s2))
}

func TestMapAndSet(t *testing.T) {
	m := map[string]string{
		"name": "snowdreams1006",
		"site": "https://snowdreams1006.github.io",
		"lang": "go",
	}

	// https://snowdreams1006.github.io
	if site, ok := m["site"]; ok {
		t.Log(site)
	} else {
		t.Log("site does not exist ")
	}

	s := map[string]bool{
		"name": true,
		"site": true,
		"lang": true,
	}

	// Pay attention to snowdreams1006
	if _, ok := m["isFollower"]; ok {
		t.Log("Have an eye on snowdreams1006")
	} else {
		s["isFollower"] = true
		t.Log("Pay attention to snowdreams1006")
	}
}

func TestMapAndSetByDelete(t *testing.T) {
	m := map[string]string{
		"name": "snowdreams1006",
		"site": "https://snowdreams1006.github.io",
		"lang": "go",
	}

	delete(m, "lang")

	// delete lang successfully
	if _,ok := m["lang"];!ok{
		t.Log("delete lang successfully")
	}
}
