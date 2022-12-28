package util_all_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(id int) {
			ret := runTask(id)
			time.Sleep(time.Millisecond * 100)
			ch <- ret
		}(i)
	}
	var finalRet = ""
	for j := 0; j < numOfRunner; j++ {
		select {
		case <-ch:
			finalRet += <-ch + "\n"
		case <-time.After(time.Millisecond * 50):
			finalRet += fmt.Sprintf("time out %d", j) + "\n"
		}
		//finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
