package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	originalSlice := make([]int, 10)

	for i := 0; i < 10; i++ {
		originalSlice[i] = rand.Intn(100)
	}

	fmt.Println(originalSlice)

	fmt.Println()
	fmt.Println(sliceExample(originalSlice))
	fmt.Println(originalSlice) //чтобы проверить не меняется ли оригинальный массив

	fmt.Println()
	fmt.Println(addElements(originalSlice, 52))
	fmt.Println(originalSlice)

	fmt.Println()
	fmt.Println(copySlice(originalSlice))
	fmt.Println(originalSlice)

	fmt.Println()
	fmt.Println(removeElement(originalSlice, 1))
	fmt.Println(originalSlice)
}

func sliceExample(slice []int) []int {
	res := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func addElements(slice []int, value int) []int {
	result := make([]int, len(slice)+1)
	copy(result, slice)
	result[len(slice)] = value
	return result
}

func copySlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	return result
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	copySlice := make([]int, len(slice))
	copy(copySlice, slice)

	return slices.Delete(copySlice, index, index+1)
}
