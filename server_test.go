package sshexec

import "testing"

func TestRun(t *testing.T) {
	s := &Server{
		Server:   "127.0.0.1:22",
		User:     "tx",
		Password: "fuckgfw",
	}
	out, err := s.Run("echo 'hello' ", "echo $HOME")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(out))
}

func TestRuns(t *testing.T) {
	s := &Server{
		Server:   "127.0.0.1:22",
		User:     "tx",
		Password: "fuckgfw",
	}
	out, err := s.Runs([]string{"echo 'hello' ", "echo $HOME"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(out))
}
