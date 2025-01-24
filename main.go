package main

import (
	"fmt"

	"github.com/tierant5/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("couldn't read the config file: %v\n", err)
	}
	err = cfg.SetUser("caleb")
	if err != nil {
		fmt.Printf("couldn't write the config file: %v\n", err)
	}
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("couldn't read the config file: %v\n", err)
	}
	fmt.Printf("DbUrl: %v\n", cfg.DbUrl)
	fmt.Printf("CurrentUserName: %v\n", cfg.CurrentUserName)
}
