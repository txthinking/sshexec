//
// cloud@txthinking.com
//
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sshexec"
	app.Version = "1.0.9"
	app.Usage = "Run command on remote server"
	app.Author = "Cloud"
	app.Email = "cloud@txthinking.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "server, s",
			Usage: "Server address, like: 1.2.3.4:22",
		},
		cli.StringFlag{
			Name:  "user, u",
			Usage: "SSH user",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "SSH password",
		},
		cli.StringSliceFlag{
			Name:  "command, c",
			Usage: "command will be run on remote server",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("server") == "" || c.String("user") == "" || c.String("password") == "" {
			cli.ShowAppHelp(c)
			return nil
		}
		if len(c.StringSlice("command")) == 0 {
			cli.ShowAppHelp(c)
			return nil
		}
		s := &Server{
			Server:   c.String("server"),
			User:     c.String("user"),
			Password: c.String("password"),
		}
		out, err := s.Runs(c.StringSlice("command"))
		if err != nil {
			return err
		}
		fmt.Print(string(out))
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
