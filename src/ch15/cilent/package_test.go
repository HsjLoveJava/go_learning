package cilent

import (
	"go_learning/src/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	//t.Log(series.square(2))
}
