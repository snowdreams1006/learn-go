package container

import (
	"fmt"
	"testing"
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

func TestSlice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// arr[2:6] =  [2 3 4 5]
	t.Log("arr[2:6] = ", arr[2:6])
	// arr[:6] =  [0 1 2 3 4 5]
	t.Log("arr[:6] = ", arr[:6])
	// arr[2:] =  [2 3 4 5 6 7 8 9]
	t.Log("arr[2:] = ", arr[2:])
	// arr[:] =  [0 1 2 3 4 5 6 7 8 9]
	t.Log("arr[:] = ", arr[:])
}

func updateSlice(s []int) {
	s[0] = 666
}

func TestUpdateSlice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

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
	// s1 =  [2 3 4 5]
	t.Log("s1 = ", s1)

	s2 := s1[3:5]
	// s2 =  [5 6]
	t.Log("s2 = ", s2)

	// index out of range
	//t.Log("s1[4] = ", s1[4])
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
	// []
	t.Log(s)

	for i := 0; i < 10; i++ {
		s = append(s, i)
	}

	// [0 1 2 3 4 5 6 7 8 9]
	t.Log(s)
}

func printSlice(s []int) {
	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))
}

func TestSliceAutoLonger(t *testing.T) {
	var s []int
	// []
	t.Log(s)

	for i := 0; i < 100; i++ {
		printSlice(s)

		s = append(s, i)
	}

	// [0 1 2 3 ...,98,99]
	t.Log(s)
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