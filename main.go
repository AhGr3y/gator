package main

import (
	"fmt"
	"log"

	"github.com/AhGr3y/gator/internal/config"
)

func main() {
	cfg := config.Config{
		DbURL: "postgres://example",
	}

	err := cfg.SetUser("Lane Wagner")
	if err != nil {
		log.Fatalf("error setting user: %v\n", err)
	}

	config, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v\n", err)
	}

	fmt.Println(config)

}
