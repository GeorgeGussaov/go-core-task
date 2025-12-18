package main

import "fmt"

func Random(cnt int) []int {
	nums := make([]int, cnt)
	ch := make(chan int)
	for i := 0; i < cnt; i++ {
		go func() {
			ch <- i
		}()
	}

	for i := 0; i < cnt; i++ {
		nums[i] = <-ch
	}
	return nums
}

func main() {
	fmt.Print(Random(15))
}
