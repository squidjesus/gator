package main

import (
	"fmt"
	"os"
	"database/sql"
	"github.com/squidjesus/gator/internal/config"
	"github.com/squidjesus/gator/internal/database"
)

import _ "github.com/lib/pq"


type state struct {
	config  *config.Config
	db		*database.Queries
}


func main() {
	// read the config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config", err)
		os.Exit(1)
	}

	// import the database
	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// initialize the 'state' struct with the cfg and db
	s := state{
		config: &cfg,
		db: database.New(db),
	}

	// creat cmds stuct and register commands
	cmds := commands{ 
		registry: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)

	// parse arguments and run cmd
	if len(os.Args) < 2 {
		fmt.Println("Expected at least 1 argument")
		os.Exit(1)
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	if err := cmds.run(&s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
