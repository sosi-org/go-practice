// package harch
package main

import "fmt"
import "time"
//import ffmt "github.com/sohale/fmt"

func reps(s string) {
    fmt.Println("hi "+s)
}

func main() {
	// print("hi")
    fmt.Println("hi")
    //var i int = 0
    //i := 0
    // make: slice (dynamic array), channels, maps.
    //var ch chan string // nil
    ch := make(chan string)
    go reps("aabbbbbce122222")
    time.Sleep(100*time.Millisecond)
    /*{
        fmt.Println(r, ch)
    }*/
}

