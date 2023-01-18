## sshexec

A command-line tool to execute remote command through ssh

### Install via [nami](https://github.com/txthinking/nami)

```
nami install sshexec
```

### Usage

```
NAME:
   sshexec - Run command on remote server

USAGE:
   sshexec [global options] command [command options] [arguments...]

VERSION:
   20230118

AUTHOR:
   Cloud <cloud@txthinking.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --server value, -s value    Server address, like: 1.2.3.4:22
   --user value, -u value      user
   --password value, -p value  password
   --key value, -k value       private key
   --command value, -c value   command will be run on remote server, ignore upload/download
   --upload value              upload file to remote server
   --download value            download file from remote server
   --to value                  dst with upload/download
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)

COPYRIGHT:
   https://www.txthinking.com
```

### Example

```
$ sshexec -s 1.2.3.4:22 -u tom -p jerry -c "ls -l"
```
