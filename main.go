package main

import (
	"fmt"
	"os"

	"github.com/tierant5/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("couldn't read the config file: %v\n", err)
		os.Exit(1)
	}

	st := &state{
		cfg: &cfg,
	}
	cmds := commands{
		cmds: map[string]func(*state, command) error{},
	}
	cmds.register("login", handerLogin)
	args := os.Args
	if len(args) < 2 {
		err = fmt.Errorf("no command found")
		fmt.Println(err)
		os.Exit(1)
	}
	cmdName := args[1]
	cmdArgs := args[2:]
	err = cmds.run(st, command{name: cmdName, args: cmdArgs})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
