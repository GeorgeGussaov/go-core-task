package main

import (
	"sync"
	"testing"
)

func Test(t *testing.T) {
	wg := NewSemaphoreWG()
	var mu sync.Mutex
	cnt := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			cnt++
			mu.Unlock()
		}()
	}
	wg.Wait()

	if cnt != 1000 {
		t.Error("wg does not work")
	}
}
