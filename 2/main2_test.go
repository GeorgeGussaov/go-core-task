package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}

	result := sliceExample(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	original := make([]int, 3, 5) //если капасити будет также 3, то даже при неправильной функции массив слайса будет пересоздан append-ом внутри функции и ошибка будет не вылавлена
	copy(original, []int{1, 2, 3})
	result := addElements(original, 4)

	if len(original) != 3 {
		t.Error("original slice was modified")
	}

	if result[len(result)-1] != 4 {
		t.Error("element was not added correctly")
	}
}
func TestCopyElements(t *testing.T) {
	original := make([]int, 3, 5)
	copy(original, []int{1, 2, 3})
	result := copySlice(original)

	if len(result) != len(original) {
		t.Error("elements was not copied")
	}
	result = append(result, 15)
	if len(original) != 3 {
		t.Error("original slice was modified")
	}
}
func TestRemoveElement(t *testing.T) {
	original := make([]int, 3, 5)
	copy(original, []int{1, 2, 3})
	_ = removeElement(original, 1)

	if original[1] != 2 {
		t.Error("original slice was modified")
	}
}
