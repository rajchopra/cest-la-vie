// Go's `sort` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.

package main

import "fmt"
import "sort"

func main() {

	// Sort methods are specific to the builtin type;
	// here's an example for strings. Note that sorting is
	// in-place, so it changes the given slice and doesn't
	// return a new one.
	strs := []string{"c", "a", "b"}
	fmt.Println("Before sorting: Strings:", strs)
	sort.Strings(strs)
	fmt.Println("\nAfter sorting: Strings:", strs)

	// check if String slice is in sorted order.
	s := sort.StringsAreSorted(strs)
	fmt.Println("String Sorted: ", s)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	fmt.Println("\n\nBefore sorting: Ints:   ", ints)
	// check if int slice is in sorted order.
	s = sort.IntsAreSorted(ints)
	fmt.Println("Ints Sorted: ", s)

	sort.Ints(ints)
	fmt.Println("\nAfter sorting: Ints:   ", ints)

	// check again if int slice is in sorted order.
	s = sort.IntsAreSorted(ints)
	fmt.Println("Ints Sorted: ", s)

}
