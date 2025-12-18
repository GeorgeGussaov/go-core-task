package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeChannels_Basic(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	go func() {
		ch3 <- 5
		ch3 <- 6
		close(ch3)
	}()

	result := make([]int, 0)
	for v := range MergeChannels(ch1, ch2, ch3) {
		result = append(result, v)
	}

	expected := []int{1, 2, 3, 4, 5, 6}

	// порядок не гарантирован → сортируем
	sort.Ints(result)
	sort.Ints(expected)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
