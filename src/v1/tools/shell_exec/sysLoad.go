package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	app := "uptime"
	printAndExecCommand(app)
}

func printAndExecCommand(app string) {

	cmd := exec.Command(app)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	load := string(stdout)
	words := strings.Fields(load)

	fmt.Println(words[5:])

	last1Min, err := strconv.ParseFloat(words[7], 64)
	last5Min, err := strconv.ParseFloat(words[8], 64)
	last15Min, err := strconv.ParseFloat(words[9], 64)

	fmt.Printf("load average over the last 1 minute: %f\n", last1Min)
	fmt.Printf("load average over the last 5 minutes: %f\n", last5Min)
	fmt.Printf("load average over the last 15 minutes: %f\n", last15Min)

}
