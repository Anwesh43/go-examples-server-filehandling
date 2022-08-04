package main 

import (
	"time"
	"fmt"
)

func handleCounter(count int, ch1 chan int, done chan int) {
	for i := 0; i <= count; i++ {
		ch1 <- i 
		time.Sleep(5 * time.Second)
	}
	done <- 0 
}	

func main() {
	ch1 := make(chan int)
	done := make(chan int)
	mainQuit := make(chan int)
	defer func() {
		close(ch1)
		close(done)
		close(mainQuit)
		fmt.Println("done with the program x-----------x hope you liked it")
	}()
	go handleCounter(10, ch1, done) 
	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("count increased to", num) 
			case <-done:
				fmt.Println("quit")
				mainQuit <- 0
				return  
			default:
				fmt.Println("waiting.....")
			}
			time.Sleep(time.Second)
		}
	}()
	<-mainQuit
	
}