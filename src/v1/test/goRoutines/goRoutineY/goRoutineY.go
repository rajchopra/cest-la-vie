// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func f(from string, duration time.Duration) {
	for i := 0; i < 200; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(duration * time.Microsecond)
	}
}

func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	f("direct", 100)

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go f("goroutine", 500)

	// start 20 goroutines of function f with different sleep duration
	for i := 0; i < 20; i++ {
		go f(strings.Join([]string{"Raj"}, strconv.Itoa(i)), time.Duration(50*(20-i)))
	}

	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` code requires we press a key
	// before the program exits.
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
