// package harch
package main

import "fmt"
//import "time"
//import ffmt "github.com/sohale/fmt"

func reps(s string, ch chan string) {
    ch <- "hi " + s
    //"hi" -> ch   error
    //var recvd string = <-ch
}

func main() {
	// print("hi")
    fmt.Println("hi")
    //var i int = 0
    //i := 0
    // make: slice (dynamic array), channels, maps.
    //var ch chan string // nil
    ch := make(chan string)
    go reps("aabbbbbce122222", ch)
    // time.Sleep(100*time.Millisecond)
    // NOT:     for s string in ch {

    //for {
        s := <-ch
        fmt.Println(s)
    //}
    /*{
        fmt.Println(r, ch)
    }*/
}

