package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const BufferSize = 50

func main() {

	var buffer [BufferSize]string
	inp := 0
	outp := 0
	count := 0
	numberOfProducts := 0
	numberOfRequests := 0
	numberOfDiscardedProducts := 0
	numberOfDiscardedRequests := 0

	c1 := make(chan string)
	c2 := make(chan int)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			fmt.Println("Product", i, "produced!")
			numberOfProducts++
			c1 <- "Product " + strconv.Itoa(i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond * 3)
		}
	}()
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			numberOfRequests++
			c2 <- 1
			fmt.Println("Consumer requested product!")
		}
	}()

	for i := 0; i < BufferSize*2; i++ {
		select {
		case s := <-c1:
			if count <= BufferSize {
				buffer[inp] = s
				fmt.Println(buffer[inp], "delivered to Buffer!")
				inp = inp + 1
				inp = (inp % BufferSize)
				count = count + 1
			} else { // discards product in s
				numberOfDiscardedProducts++
				fmt.Println(s, "discarded by Buffer!")
			}

		case <-c2:
			if count > 0 {
				fmt.Println(buffer[outp], " sent to Consumer!")
				outp = outp + 1
				outp = (outp % BufferSize)
				count = count - 1
			} else { //discards requests
				numberOfDiscardedRequests++
				fmt.Println("Request discarded by Buffer!")
			}
		}
	}
	fmt.Println("\n\nReport\n")
	fmt.Println("Number of Products", numberOfProducts)
	fmt.Println("Number of Requests", numberOfRequests)
	fmt.Println("Number of Discarded Products", numberOfDiscardedProducts)
	fmt.Println("Number of Discarded Requests", numberOfDiscardedRequests)
	fmt.Println("\n\nList of", count, " products in Buffer\n")
	for i := 0; i < count; i++ {
		fmt.Println(buffer[outp])
		outp = outp + 1
		outp = (outp % BufferSize)
	}
	fmt.Println("Thank Producer and Consumer! You've done a wonderful job! I'm leaving.")

}
