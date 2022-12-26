package err_test

import (
	"errors"
	"testing"
)

var LessThanTwoError = errors.New("it is less than 2")

var LagerThenHundredError = errors.New("it is lager than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LagerThenHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, error := GetFibonacci(1); error != nil {
		if error == LessThanTwoError {
			t.Log("it is less.")
		}
		t.Error(error)
	} else {
		t.Log(v)
	}
}
