package main

import (
	"fmt"
	"os/exec"
)

func main() {
	app := "/sbin/ifconfig"
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
	fmt.Println(load)

}
