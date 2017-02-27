## sshexec

A command-line tool to execute remote command through ssh

### Install
```
$ go get github.com/txthinking/sshexec/cli/sshexec
```

### Usage

```
$ sshexec -h
NAME:
   sshexec - Run command on remote server

USAGE:
   sshexec [global options] command [command options] [arguments...]

VERSION:
   1.0.9

AUTHOR:
   Cloud <cloud@txthinking.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --server value, -s value    Server address, like: 1.2.3.4:22
   --user value, -u value      SSH user
   --password value, -p value  SSH password
   --command value, -c value   command will be run on remote server
   --help, -h                  show help
   --version, -v               print the version
```

### Example
```
$ sshexec -s 1.2.3.4:22 -u tx -p mypassword -c "ls -l" -c "echo hello" -c "echo world"
```
