package main

import "fmt"
import "time"

var countryTz = map[string]string{
	"Hungary": "Europe/Budapest",
	"Egypt":   "Africa/Cairo",
}

func timeIn(name string) time.Time {
	loc, err := time.LoadLocation(countryTz[name])
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}

func main() {
	utc := time.Now().UTC().Format("15:04")
	hun := timeIn("Hungary").Format("15:04")
	eg := timeIn("Egypt").Format("15:04")
	fmt.Println(utc, hun, eg)
}
