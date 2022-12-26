package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}

	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	var arr = [...]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		t.Log(arr[i])
	}

	for _, e := range arr {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	arrSec := arr[:3]
	t.Log(arrSec)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	//if a == b {
	//	t.Log("equal")
	//}
	a = nil
	b = nil
	if a == nil && b == nil {
		t.Log("a is nil and b is nil")
	}
}
