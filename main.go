package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"time"

	"strings"

	. "gossh/depends/exception"

	. "gossh/depends/console"
	ssh "gossh/depends/sshclient"

	"github.com/zh-five/golimit"
)

func ticks() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func parseInt(expr string, defval int) int {
	result := defval
	Try(func() {
		val, err := strconv.Atoi(expr)
		if nil != err {
			Throw(err.Error())
		}
		result = val
	}, func(err error) {
		LogErrorln(err.Error())
	})
	return result
}
func main() {
	// ./gossh -c "hostname" -p unicloud -i "127.0.0.1 127.0.0.1"
	// ./gossh -c "hostname" -p unicloud -i "`seq 250 254|xargs -I[] echo '10.253.27.'[]|xargs`"
	// ./gossh -c "ping -w2 -c1 10.253.27.250" -p unicloud -i "`seq 10 50|xargs -I[] echo '10.253.27.'[]|xargs`"

	ips := ""
	flag.StringVar(&ips, "i", "", "ipaddress\n\texample:\n\t\"192.168.0.1 192.168.0.1\"\n\t\"root:123456@192.168.0.1:23 admin@192.168.0.1 192.168.0.1\" ")
	port := 22
	flag.IntVar(&port, "O", 22, "host port")
	timeout := 5
	flag.IntVar(&timeout, "t", 5, "connect timeout")
	user := ""
	flag.StringVar(&user, "u", "root", "user")
	pass := ""
	flag.StringVar(&pass, "p", "", "password")
	cmd := ""
	flag.StringVar(&cmd, "c", "", "command")
	debug := false
	flag.BoolVar(&debug, "d", false, "debug")
	parallel := uint(64)
	flag.UintVar(&parallel, "n", 64, "parallel number")

	flag.Parse()
	if len(strings.Trim(ips, " ")) < 1 {
		LogErrorln("ipaddress cannot be empty")
		return
	}
	if len(cmd) < 1 {
		LogErrorln("command cannot be empty")
		return
	}
	addrs := strings.Split(ips, " ")
	type Task struct {
		Host    string
		Elapsed string
		Index   int
		Result  []string
		Error   string
	}
	results := make(chan Task, len(addrs))
	start := time.Now()
	if debug {
		LogInfoln(ticks(), " starting...")
		LogInfoln(ticks(), fmt.Sprintf(" %s#%s@host:%d<\"%s\"", user, pass, port, cmd))
	}
	g := golimit.NewGoLimit(parallel)
	for i, addr := range addrs {
		g.Add()
		_addr := addr
		_port := port
		_user := user
		_pass := pass
		if -1 != strings.IndexAny(addr, "@") {
			parts := strings.Split(addr, "@")
			uandp := parts[0]
			_user = uandp
			if -1 != strings.IndexAny(uandp, ":") {
				userpass := strings.Split(uandp, ":")
				_user = userpass[0]
				_pass = userpass[1]
			}
			iandp := parts[1]
			_addr = iandp
			if -1 != strings.IndexAny(iandp, ":") {
				ipport := strings.Split(iandp, ":")
				_addr = ipport[0]
				_port = parseInt(ipport[1], 22)
			}
		}
		go func(g *golimit.GoLimit, i int, addr string, port int, user string, pass string, cmd string, timeout int) {
			defer g.Done()
			begin := time.Now()
			result, err := ssh.Exec(addr, port, user, pass, cmd, timeout)
			msg := ""
			if nil != err {
				msg = fmt.Sprintf("%v", err)
			}
			task := Task{
				Host:    fmt.Sprintf("%s:%s@%s:%d", user, pass, addr, port),
				Elapsed: time.Since(begin).String(),
				Index:   i,
				Result:  strings.Split(result, "\n"),
				Error:   msg,
			}
			results <- task
		}(g, i, _addr, _port, _user, _pass, cmd, timeout)
	}
	if debug {
		LogInfoln(ticks(), " tasks added")
	}
	g.WaitZero()
	tasks := struct {
		Tasks   []Task
		Count   int
		Success bool
		Elapsed string
	}{
		Tasks:   make([]Task, 0),
		Count:   len(addrs),
		Elapsed: "",
		Success: true,
	}

	for i := 0; i < len(addrs); i++ {
		v := <-results
		tasks.Tasks = append(tasks.Tasks, v)
		if len(v.Error) > 0 {
			tasks.Success = false
		}
	}
	tasks.Elapsed = time.Since(start).String()
	j, _ := json.MarshalIndent(tasks, "", "\t")
	Logln(string(j))
	if debug {
		LogInfoln(ticks(), " All host (", len(addrs), ") tasks have been executed successfully, elapsed:", time.Since(start))
	}
}
