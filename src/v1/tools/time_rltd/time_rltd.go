package main

import "fmt"
import "time"

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC850))
	fmt.Println(t.Format("20060102150405"))
	fmt.Println(t.Format("2006-01-02-15:04:05"))
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("2006-01-02<raj>15:04:05"))

}
