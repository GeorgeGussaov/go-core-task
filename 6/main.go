package main

import "fmt"

//Как я понял нужно возвращать слайс случайных чисел. Плюс длина слайса равна диапазону рандомных значений, надеюсь некритично
func Random(cnt int) []int {
	nums := make([]int, cnt)
	ch := make(chan int)
	for i := 0; i < cnt; i++ {
		go func(val int) {
			ch <- i
		}(i)
	}

	for i := 0; i < cnt; i++ {
		nums[i] = <-ch
	}
	return nums
}

func main() {
	fmt.Print(Random(15))
}
