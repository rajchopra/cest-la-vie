package main

import "os/exec"

func main() {
	app := "echo"
	arg := []string{"-e", "Hi Raj!!"}
	printAndExecCommand(app, arg...)
}

func printAndExecCommand(app string, args ...string) {

	cmd := exec.Command(app, args...)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(stdout))

}
