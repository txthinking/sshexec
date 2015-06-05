## sshexec

A command-line tool to execute remote command through ssh

### Install
```
$ go get github.com/txthinking/toolbox/sshexec
```

### Usage

0. create config `config.json`:

```
{
    "web01": {
        "Server": {
            "IP": "192.168.1.9",
            "Port": 22,
            "User": "tx",
            "Password": "fuckgfw"
        },
        "Commands": [
            "cd /tmp/",
            "ls -l"
        ]
    },
    "web02": {
        "Server": {
            "IP": "192.168.1.10",
            "Port": 22,
            "User": "tx",
            "Password": "fuckgfw"
        },
        "Commands": [
            "ls -l /home"
        ]
    }
}
```

1. Usage

```
$ sshexec -h
```
