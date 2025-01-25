package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/tierant5/gator/internal/config"
	"github.com/tierant5/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("couldn't read the config file: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Printf("could not open connection to database: %v", err)
	}

	dbQueries := database.New(db)

	st := &state{
		cfg: &cfg,
		db:  dbQueries,
	}
	cmds := commands{
		cmds: map[string]func(*state, command) error{},
	}

	cmds.register("login", handerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

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
		log.Fatal(err)
	}
}
