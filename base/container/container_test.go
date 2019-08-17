package container

import "testing"

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
