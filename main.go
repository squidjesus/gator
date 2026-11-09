package main

import (
	"fmt"

	"github.com/squidjesus/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config", err)
		return
	}
	cfg.SetUser("jsquid")
	cfg, err = config.Read()
	if err != nil {
		fmt.Println("error reading config", err)
	}
	fmt.Printf("%+v\n", cfg)
}
