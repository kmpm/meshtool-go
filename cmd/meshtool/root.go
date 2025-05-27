package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/urfave/cli/v3"
)

var (
	once sync.Once
	root *cli.Command
)

func getRoot() *cli.Command {
	once.Do(func() {
		root = &cli.Command{
			Name:  "meshtool",
			Usage: "Meshtool CLI",
			Action: func(ctx context.Context, cmd *cli.Command) error {
				fmt.Println("boom")
				return nil
			},
			Commands: []*cli.Command{},
		}
	})
	return root
}
