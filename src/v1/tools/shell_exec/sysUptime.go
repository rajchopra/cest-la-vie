package main

import "os/exec"

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

	print(string(stdout))

}
