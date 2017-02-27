package sshexec

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

type Server struct {
	Server   string
	User     string
	Password string
}

func (s *Server) Run(cmd ...string) ([]byte, error) {
	return s.Runs(cmd)
}

func (s *Server) Runs(cmd []string) ([]byte, error) {
	keyboardInteractiveChallenge := func(
		user,
		instruction string,
		questions []string,
		echos []bool,
	) (answers []string, err error) {
		if len(questions) == 0 {
			return []string{}, nil
		}
		return []string{s.Password}, nil
	}
	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(keyboardInteractiveChallenge),
			ssh.Password(s.Password),
		},
	}
	conn, err := ssh.Dial("tcp", s.Server, config)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := strings.Join(cmd, " && ")
	output, err := session.CombinedOutput(fmt.Sprintf("sh -c '%s'", c))
	if err != nil {
		return nil, err
	}
	return output, err
}
