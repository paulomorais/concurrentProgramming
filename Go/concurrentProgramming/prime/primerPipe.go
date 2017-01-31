package prime

func Generate(max int, ch chan<- int) {
	for i := 2; i <= max; i++ {
		ch <- i
	}
	ch <- 0
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
		if i == 0 {
			out <- i
			break
		}
	}
}

func RunTest(max int) {
	ch := make(chan int)
	go Generate(max, ch)
	for {
		prime := <-ch
		if prime == 0 {
			close(ch)
			break
		}
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}