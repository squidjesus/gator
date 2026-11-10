package main

import (
	"fmt"
	"os"
	"github.com/squidjesus/gator/internal/config"
)

type state struct {
	config *config.Config
}


func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config", err)
		return
	}
	s := state{ &cfg }
	cmds := commands{ 
		registry: map[string]func(*state, command) error{},
	}

	// register commands
	cmds.register("login", handlerLogin)

	// parse arguments and run cmd
	if len(os.Args) < 2 {
		fmt.Println("Expected at least 1 argument")
		return
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	if err := cmds.run(&s, cmd); err != nil {
		fmt.Println(err)
		return
	}
}
