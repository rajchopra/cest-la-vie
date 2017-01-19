// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

package main

import "fmt"
import "os"

type gatsby struct {
	valOne, valTwo int
}

func main() {

	// Go offers several printing "verbs" designed to
	// format general Go values. For example, this prints
	// an instance of our `gatsby` struct.
	p := gatsby{22, 11}
	fmt.Printf("%v\n", p)
	// OUTPUT : {22 11}

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.
	fmt.Printf("%+v\n", p)
	// OUTPUT : {valOne:22 valTwo:11}

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.
	fmt.Printf("%#v\n", p)
	// OUTPUT : main.gatsby{valOne: 22, valTwo: 11}

	// To print the type of a value, use `%T`.
	fmt.Printf("%T\n", p)
	// OUTPUT : main.gatsby

	// Formatting booleans is straight-forward.
	fmt.Printf("%t\n", true)
	// OUTPUT : true

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.
	fmt.Printf("%d\n", 123)
	// OUTPUT : 123

	// This prints a binary representation.
	fmt.Printf("%b\n", 14)
	// OUTPUT : 1110

	// This prints the character corresponding to the
	// given integer.
	fmt.Printf("%c\n", 33)
	// OUTPUT : !

	// `%x` provides hex encoding.
	fmt.Printf("%x\n", 456)
	// OUTPUT : 1c8

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.
	fmt.Printf("%f\n", 78.9)
	// OUTPUT : 78.900000

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.
	fmt.Printf("%e\n", 123400000.0)
	// OUTPUT : 1.234000e+08
	fmt.Printf("%E\n", 123400000.0)
	// OUTPUT : 1.234000E+08

	// For basic string printing use `%s`.
	fmt.Printf("%s\n", "Raj Vimal Chopra")
	// OUTPUT : Raj Vimal Chopra

	// For basic string with double-quotes use `%s`.
	fmt.Printf("%s\n", "\"TheGreatGatsby\"")
	// OUTPUT : "TheGreatGatsby"

	// To double-quote strings as in Go source, use `%q`.
	fmt.Printf("%q\n", "\"TheGreatGatsby\"")
	// OUTPUT : "\"TheGreatGatsby\""

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.
	fmt.Printf("%x\n", "hex this")
	// OUTPUT : 6865782074686973

	// To print a representation of a pointer, use `%p`.
	fmt.Printf("%p\n", &p)
	// OUTPUT : 0xc82000a2c0

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// OUTPUT : |    12|   345|

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// OUTPUT : |  1.20|  3.45|

	// To left-justify, use the `-` flag.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// OUTPUT : |1.20  |3.45  |

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.
	fmt.Printf("|%9s|%9s|\n", "Raj", "Chopra")
	// OUTPUT : |      Raj|   Chopra|

	// To left-justify use the `-` flag as with numbers.
	fmt.Printf("|%-9s|%-9s|\n", "Raj", "Chopra")
	// OUTPUT : |Raj      |Chopra   |

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.
	s := fmt.Sprintf("The %s", "Chopras")
	fmt.Println(s)
	// OUTPUT : The Chopras

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.
	fmt.Fprintf(os.Stderr, "The %s\n", "GreatGatsby")
	// OUTPUT : The GreatGatsby

}
