package main

import "fmt"

func Conveyor(ch1 <-chan uint8, ch2 chan<- float64) {
	for num := range ch1 {
		f := float64(num)
		ch2 <- f * f * f
	}
	close(ch2)
}

func main() {
	ch1 := make(chan uint8)
	ch2 := make(chan float64)

	go Conveyor(ch1, ch2)

	go func() {
		for _, v := range []uint8{1, 2, 3, 4, 5} {
			ch1 <- v
		}
		close(ch1)
	}()

	for res := range ch2 {
		fmt.Println(res)
	}
}
