package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	app := "hostname"
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

	fmt.Println(words[0:])

}
