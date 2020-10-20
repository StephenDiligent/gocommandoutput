package main

import (
	"fmt"
	"log"

	"github.com/StephenDiligent/gocommandoutput/internal/exec"
)

type command struct {
	name  string
	args  []string
	input string
}

func main() {
	commands := make(map[int]command)
	commands[1] = command{name: "cmd", args: []string{"/C", "echo", "helloworld"}, input: "helloworld"}
	commands[2] = command{name: "powershell", args: []string{"echo", "helloworld"}, input: "helloworld"}
	commands[3] = command{name: "more", args: []string{"helloworld.txt"}, input: "helloworld"}
	commands[4] = command{name: "cmd", args: []string{"/C", "more", "helloworld.txt"}, input: "helloworld"}

	commands[5] = command{name: "cmd", args: []string{"/C", "echo", "hello world"}, input: "hello world"}
	commands[6] = command{name: "powershell", args: []string{"echo", "hello world"}, input: "hello world"}
	commands[7] = command{name: "more", args: []string{"hello world.txt"}, input: "hello world"}
	commands[8] = command{name: "cmd", args: []string{"/C", "more", "hello world.txt"}, input: "hello world"}

	for id, command := range commands {
		res, err := exec.Command(command.name, command.args...).Run()
		if err != nil {
			log.Fatal(err)
		}

		output := res.StrOutput()

		if output != command.input {
			fmt.Printf("command %d output [%s] doesn't match input [%s]\n", id, output, command.input)
		} else {
			fmt.Printf("command %d output [%s] matches input [%s]\n", id, output, command.input)
		}
	}

}
