package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go Conveyor(in, out)

	go func() {
		in <- 1
		in <- 2
		in <- 3
		close(in)
	}()

	var result []float64
	for v := range out {
		result = append(result, v)
	}

	expected := []float64{1, 8, 27}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
