package main

import (
	"fmt"
	"time"
)

// func pro1() {
// 	time.Sleep(100 * time.Millisecond)
// 	fmt.Println("Executing function1")
// }
// func pro2() {
// 	time.Sleep(100 * time.Millisecond)
// 	fmt.Println("Executing function2")
// }
// func pro3() {
// 	time.Sleep(100 * time.Millisecond
// 	fmt.Println("Executing function3")
// }
// func main() {
// 	now := time.Now()
// 	go pro1()
// 	go pro2()
// 	go pro3()
// 	time.Sleep(100 * time.Millisecond)
// 	fmt.Println("Time elapsed", time.Since(now))
// }

//using waitgroups
// func main() {
// 	now:= time.Now()
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() { //fork point
// 		defer wg.Done()
// 		work()
// 	}()
// 	wg.Wait()
// 	fmt.Println("Time elapsed: ",time.Since(now))
// 	fmt.Println("Done waiting, main exists")
// 	fmt.Println("Time elapsed: ",time.Since(now))

// }

// func work() {
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Println("...Printing dome stuff")
// }

// using channels
// func main() {
// 	now := time.Now()
// 	done := make(chan struct{})
// 	go func() { //fork point
// 		work()
// 		done <- struct{}{}
// 	}()
// 	<-done //join point
// 	fmt.Println("Time elapsed: ", time.Since(now))
// 	fmt.Println("Done waiting, main exists")

// }

// func work() {
// 	time.Sleep(500 * time.Millisecond)
// 	fmt.Println("...Printing dome stuff")
// }

func pro1(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Executing function1")
	done <- struct{}{} // Signal that we are done.
}
func pro2(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Executing function2")
	done <- struct{}{} // Signal that we are done.
}
func pro3(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Executing function3")
	done <- struct{}{} // Signal that we are done.
}
func main() {
	now := time.Now()
	done := make(chan struct{})
	go pro1(done)
	go pro2(done)
	go pro3(done)
	<-done //waits for all goroutines to finish

	// <-done is used to block the main goroutine until one signal is received on the done channel. This means that the main goroutine will wait until at least one of the pro1, pro2, or pro3 functions has completed its work and signaled its completion.
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Time elapsed", time.Since(now))
}
