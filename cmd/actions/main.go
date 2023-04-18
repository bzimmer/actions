package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/bzimmer/actions"
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
		fmt.Fprintf(app.Writer, "%v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
