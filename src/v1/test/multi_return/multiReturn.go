// Go has built-in support for _multiple return values_.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.

package main

import "fmt"

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func vals() (int, int) {
	return 3, 7
}

// SumProductDiff : the function returns 3 `int`s.
func SumProductDiff(i, j int) (int, int, int) {
	return i + j, i * j, i - j
}

// max : the function returns 2 `int`s.
// Max number as first int.
// Min number as second int.
func max(num1, num2 int) (int, int) {
	if num1 > num2 {
		return num1, num2
	} else {
		return num2, num1
	}
}

func main() {

	// Here we use the 2 different return values from the
	// call with _multiple assignment_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values,
	// use the blank identifier `_`.
	_, c := vals()
	fmt.Println(c)

	sum, prod, diff := SumProductDiff(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)

	var a1 int = 150
	var a2 int = 130
	func(i ...int) { fmt.Printf("Max val: %d Min val: %d\n", i[0], i[1]) }(max(a1, a2))

}
