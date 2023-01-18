// Copyright (c) 2017-present Cloud <cloud@txthinking.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
	"github.com/txthinking/hancock"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "sshexec"
	app.Version = "20230118"
	app.Usage = "Run command on remote server"
	app.Authors = []*cli.Author{
		{
			Name:  "Cloud",
			Email: "cloud@txthinking.com",
		},
	}
	app.Copyright = "https://www.txthinking.com"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Server address, like: 1.2.3.4:22",
		},
		&cli.StringFlag{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "user",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "password",
		},
		&cli.StringFlag{
			Name:    "key",
			Aliases: []string{"k"},
			Usage:   "private key",
		},
		&cli.StringFlag{
			Name:    "command",
			Aliases: []string{"c"},
			Usage:   "command will be run on remote server, ignore upload/download",
		},
		&cli.StringFlag{
			Name:  "upload",
			Usage: "upload file to remote server",
		},
		&cli.StringFlag{
			Name:  "download",
			Usage: "download file from remote server",
		},
		&cli.StringFlag{
			Name:  "to",
			Usage: "dst with upload/download",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("server") == "" || c.String("user") == "" {
			cli.ShowAppHelp(c)
			return nil
		}
		if c.String("command") == "" && c.String("upload") == "" && c.String("download") == "" {
			cli.ShowAppHelp(c)
			return nil
		}
		if (c.String("upload") != "" || c.String("download") != "") && c.String("to") == "" {
			cli.ShowAppHelp(c)
			return nil
		}
		var b []byte
		if c.String("key") != "" {
			var err error
			b, err = os.ReadFile(c.String("key"))
			if err != nil {
				return err
			}
		}
		i, err := hancock.NewInstance(c.String("server"), c.String("user"), c.String("password"), b)
		if err != nil {
			return err
		}
		defer i.Client.Close()
		if c.String("command") != "" {
			s, err := i.Client.NewSession()
			if err != nil {
				return err
			}
			defer s.Close()
			s.Stdout = os.Stdout
			s.Stderr = os.Stderr
			err = s.Run(c.String("command"))
			if err != nil {
				return err
			}
			return nil
		}
		if c.String("upload") != "" {
			src, err := os.Open(c.String("upload"))
			if err != nil {
				return err
			}
			defer src.Close()
			sc, err := sftp.NewClient(i.Client)
			if err != nil {
				return err
			}
			defer sc.Close()
			dst, err := sc.Create(c.String("to"))
			if err != nil {
				return err
			}
			defer dst.Close()
			if _, err := io.Copy(dst, src); err != nil {
				return err
			}
			return nil
		}
		if c.String("download") != "" {
			sc, err := sftp.NewClient(i.Client)
			if err != nil {
				return err
			}
			defer sc.Close()
			src, err := sc.Open(c.String("download"))
			if err != nil {
				return err
			}
			defer src.Close()
			dst, err := os.Create(c.String("to"))
			if err != nil {
				return err
			}
			defer dst.Close()
			if _, err := io.Copy(dst, src); err != nil {
				return err
			}
			return nil
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Println(err)
	}
}
