package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	expected := []int{64, 3}

	ok, result := GetCommonValues(a, b)

	if !reflect.DeepEqual(result, expected) && ok == true {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
