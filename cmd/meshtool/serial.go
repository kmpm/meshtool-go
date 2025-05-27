package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/kmpm/meshtool-go/public/transport/serial"
	"github.com/urfave/cli/v3"
)

func init() {
	getRoot().Commands = append(getRoot().Commands,
		&cli.Command{
			Name: "serial",
			Commands: []*cli.Command{
				{
					Name:  "list",
					Usage: "List all serial ports",
					Action: func(ctx context.Context, cmd *cli.Command) error {
						ports := serial.GetPorts()
						log.Info("Serial ports found:", "ports", len(ports))
						for _, port := range ports {
							fmt.Println(port)
						}
						return nil
					},
				},
			},
		})

}
