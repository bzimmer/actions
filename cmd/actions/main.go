package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/bzimmer/actions"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "actions",
		HelpName: "actions",
		Usage:    "simple command line tool for testing",
		Action: func(c *cli.Context) error {
			enc := json.NewEncoder(c.App.Writer)
			return enc.Encode(actions.Support())
		},
	}
	if err := app.RunContext(context.Background(), os.Args); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
