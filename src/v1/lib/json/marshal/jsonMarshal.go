// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import "encoding/json"
import "fmt"

// Skyline1 : We'll use these two structs to demonstrate encoding and
// decoding of custom types below.
type Skyline1 struct {
	Page int
	Cars []string
}

// Skyline2 : Similar to struct Skyline1, but with JSON return Key name.
type Skyline2 struct {
	Page int      `json:"page"`
	Cars []string `json:"cars"`
}

func main() {

	// First we'll look at encoding basic data types to
	// JSON strings. Here are some examples for atomic
	// values.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher raj")
	fmt.Println(string(strB))

	// And here are some for slices and maps, which encode
	// to JSON arrays and objects as you'd expect.
	slcD := []string{"BMW", "AUDI", "MERC"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"BMW": 5, "AUDI": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The JSON package can automatically encode your
	// custom data types. It will only include exported
	// fields in the encoded output and will by default
	// use those names as the JSON keys.
	starwars1D := &Skyline1{
		Page: 1,
		Cars: []string{"BMW", "AUDI", "MERC"}}
	starwars1B, _ := json.Marshal(starwars1D)
	fmt.Println(string(starwars1B))

	// You can use tags on struct field declarations
	// to customize the encoded JSON key names. Check the
	// definition of `Skyline2` above to see an example
	// of such tags.
	starwars2D := &Skyline2{
		Page: 1,
		Cars: []string{"BMW", "AUDI", "MERC"}}
	starwars2B, _ := json.Marshal(starwars2D)
	fmt.Println(string(starwars2B))
}
