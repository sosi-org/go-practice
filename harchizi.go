// package harch
package main

import "fmt"

import "time"

//import ffmt "github.com/sohale/fmt"

func reps(s string, ch chan string) {
	fmt.Println("1")
	ch <- "hi " + s
	// Blocking: buffer is ZERO. As if YIELD.
	fmt.Println("2")
	// Blocking
	ch <- "bye " + s
	fmt.Println("3")
	//"hi" -> ch   error
	//var recvd string = <-ch
}

func main() {
	fmt.Println("hi")
	//var i int = 0
	//i := 0
	// make: slice (dynamic array), channels, maps.
	//var ch chan string // nil
	ch := make(chan string)
	var ch2 chan string = make(chan string)
	go reps("X from ch", ch)
	go func(s string, ch2 chan string) {
		ch2 <- "hi " + s
	}("Y from ch2", ch2)

	// NOT:     for s string in ch {

	//for {
	//}
	s := <-ch2
	fmt.Println(s)

	s = <-ch
	fmt.Println(s)
	fmt.Println(<-ch)
	// channels are for coordination. synchronisation.

	/*
	   var a string
	   a = <-ch  receive from channel
	   ch <-a  send to channel

	*/
	time.Sleep(100 * time.Millisecond)
}

/*
go run harchizi.go
*/
