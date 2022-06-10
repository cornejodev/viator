package main

import (
	"fmt"
	"os"

	"github.com/cornejodev/viator/config"
	"github.com/cornejodev/viator/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := app.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
