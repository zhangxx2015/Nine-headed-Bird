package SSHClient

import (
	"bytes"
	"net"
	"time"

	"fmt"

	. "gossh/depends/exception"

	"golang.org/x/crypto/ssh"
)

type Conn struct {
	net.Conn
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (c *Conn) Read(b []byte) (int, error) {
	err := c.Conn.SetReadDeadline(time.Now().Add(c.ReadTimeout))
	if err != nil {
		return 0, err
	}
	return c.Conn.Read(b)
}

func (c *Conn) Write(b []byte) (int, error) {
	err := c.Conn.SetWriteDeadline(time.Now().Add(c.WriteTimeout))
	if err != nil {
		return 0, err
	}
	return c.Conn.Write(b)
}

func dialTimeout(network, addr string, config *ssh.ClientConfig, timeout time.Duration) (*ssh.Client, error) {

	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return nil, err
	}

	timeoutConn := &Conn{conn, timeout, timeout}
	c, chans, reqs, err := ssh.NewClientConn(timeoutConn, addr, config)
	if err != nil {
		return nil, err
	}
	client := ssh.NewClient(c, chans, reqs)

	// this sends keepalive packets every 2 seconds
	// there's no useful response from these, so we can just abort if there's an error
	go func() {
		t := time.NewTicker(2 * time.Second)
		defer t.Stop()
		for range t.C {
			_, _, err := client.Conn.SendRequest("keepalive@golang.org", true, nil)
			if err != nil {
				return
			}
		}
	}()
	return client, nil
}

func Exec(addr string, port int, user string, pass string, cmd string, timeout int) (string, error) {
	result := ""
	var err error = nil
	Try(func() {
		client, err := dialTimeout("tcp", fmt.Sprintf("%s:%d", addr, port), &ssh.ClientConfig{
			User:            user,
			Auth:            []ssh.AuthMethod{ssh.Password(pass)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}, time.Second*time.Duration(timeout))
		if err != nil {
			Throw(fmt.Sprintf("SSH dial error: %s", err.Error()))
		}
		defer client.Close()
		session, err := client.NewSession()
		if err != nil {
			Throw(fmt.Sprintf("new session error: %s", err.Error()))
		}
		defer session.Close()
		var b bytes.Buffer
		session.Stdout = &b
		if err := session.Run(cmd); err != nil {
			Throw("Failed to run: " + err.Error())
		}
		result = b.String()

	}, func(ex error) {
		// println(err.Error())
		// println("catch")
		err = ex
	})
	return result, err
}
