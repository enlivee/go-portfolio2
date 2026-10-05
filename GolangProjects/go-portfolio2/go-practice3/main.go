package main

// func Send(ch chan int, num int) {
// 	ch <- num
// }

func Receive(ch chan int) int {
	val := <- ch
	return val
}

func Send(ch1, ch2 chan int) {
	go func() {
		for i := 0; i <= 2; i++ {
		ch1 <- i
	}}()
	go func() {
		for i := 0; i <= 2; i++ {
		ch2 <- i
	}
	}()
}

func Process(nums []int) chan int {
	ch := make(chan int, 10)
	for _, num := range nums {
		ch <- num
	}
	return ch
}