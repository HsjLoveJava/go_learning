package map_ext

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return op * op * op }
	t.Log(m1[1](2), m1[2](2), m1[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = false
	mySet[3] = true
	t.Log(len(mySet))
	deleteMapValue(mySet, 3)
	n := 3
	if v, ok := mySet[n]; ok {
		t.Logf("%d is existing is %t.", n, v)
	} else {
		t.Logf("%d is not existing.", n)
	}
	t.Log(len(mySet))
}

func deleteMapValue(myMap map[int]bool, key int) {
	delete(myMap, key)
}
