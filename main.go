package main

import (
	"fmt"

	"github.com/AhGr3y/gator/internal/config"
)

func main() {
	cfg := config.Config{
		DbURL: "postgres://example",
	}

	err := cfg.SetUser("Johnny Duo")
	if err != nil {
		fmt.Printf("error setting user: %v\n", err)
	}

	config, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
	}

	fmt.Println(config)

}
