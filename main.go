//
// cloud@txthinking.com
//
package main

import(
    "golang.org/x/crypto/ssh"
    "fmt"
    "bytes"
    "strings"
    "strconv"
    "os"
    "io/ioutil"
    "encoding/json"
    "flag"
	"log"
)

type Server struct{
    IP string
    Port int
    User string
    Password string
}

type One struct{
    Server *Server
    Commands []string
}

type All map[string]*One

func Usage(){
        var usage string = `A Simple Publish System
Usage:
    -h      Help.
    -c      The path of config file. Default ./config.json.
    -l      Show list of all environments.
    -a      Run all.
    -w      Which one will be connected and run. If more like "web01:web02".
    -args   This arguments can be passed in config.json.
            Like "tom" or more "tom:jerry".
            Used like ${1}, ${2}

Creator: Cloud <cloud@txthinking.com>
`
        fmt.Print(usage)
}

func GetConfig(configFile string, args []string) (all All, err error){
    var file *os.File
    var data []byte
    file, err = os.OpenFile(configFile, os.O_RDONLY, 0444)
    if err != nil {
        return
    }
    data, err = ioutil.ReadAll(file)
    if err != nil {
        return
    }
    var i int
    var s string
    for i, s = range args{
        data = bytes.Replace(data, []byte(fmt.Sprintf("${%d}", i+1)), []byte(s), -1)
    }
    err = json.Unmarshal(data, &all)
    if err != nil {
        return
    }
    return
}

func Run(s *Server, commands []string)(output string, err error){
    var config *ssh.ClientConfig
    var conn *ssh.Client
    var session *ssh.Session

    config = &ssh.ClientConfig{
        User: s.User,
        Auth: []ssh.AuthMethod{
            ssh.Password(s.Password),
        },
    }
    conn, err = ssh.Dial("tcp", s.IP+":"+strconv.Itoa(s.Port), config)
    if err != nil {
        return
    }
    defer conn.Close()

    session, err = conn.NewSession()
    if err != nil {
        return
    }
    defer session.Close()

    var o bytes.Buffer
    session.Stdout = &o
    var c string = strings.Join(commands, " 2>&1 && ")
    if err = session.Run(fmt.Sprintf("sh -c '%s'", c)); err != nil{
        return
    }
    output = o.String()
    return
}

var h bool
var l bool
var a bool
var w string
var c string
var args string

func main(){
    var err error
    var all All
    var which string
    var one *One
    var output string
    var ok bool

    flag.BoolVar(&h, "h", false, "Usage.")
    flag.BoolVar(&l, "l", false, "List of the environments.")
    flag.BoolVar(&a, "a", false, "Run all.")
    flag.StringVar(&w, "w", "", "Which will be connected and run.")
    flag.StringVar(&c, "c", "config.json", "Path of config file.")
    flag.StringVar(&args, "args", "", "Arguments will be passed in config file.")
    flag.Parse()
    if h {
        Usage()
        return
    }

    // read config
    if c == ""{
        c = "config.json"
    }

    var argss []string
    var s string
    argss = make([]string, 0)
    for _, s = range strings.Split(args, ":"){
        s = strings.TrimSpace(s)
        if s != ""{
            argss = append(argss, s)
        }
    }

    all, err = GetConfig(c, argss)
    if err != nil{
        log.Fatal(err)
        return
    }

    // show list
    if l {
        for which, _ = range all{
            fmt.Println(which)
        }
        return
    }

    // run all
    if a {
        for which, one = range all{
            output, err = Run(one.Server, one.Commands)
            if err != nil{
                log.Printf("[%s] %s\n", which, err.Error())
                continue
            }
            log.Printf("[%s]\n%s\n", which, output)
        }
        return
    }

    if w == ""{
        Usage()
        return
    }

    // run some
    for _, which = range strings.Split(w,":"){
        which = strings.TrimSpace(which)
        if which == ""{
            continue
        }
        one, ok = all[which]
        if !ok{
            log.Printf("[%s] %s\n", which, "not found")
            continue
        }
        output, err = Run(one.Server, one.Commands)
        if err != nil{
            log.Printf("[%s] %s\n", which, err.Error())
            continue
        }
        log.Printf("[%s]\n%s\n", which, output)
    }

}

