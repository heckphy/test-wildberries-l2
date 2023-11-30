package telnet

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type Client struct {
	conn net.Conn
}

func NewClient(host string, port string, timeout time.Duration) (*Client, error) {
	address := net.JoinHostPort(host, port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Start() {
	defer c.conn.Close()

	fmt.Println("Connected to", c.conn.RemoteAddr())

	go c.read()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading user input:", err)
			continue
		}

		_, err = c.conn.Write([]byte(line))
		if err != nil {
			fmt.Println("Error sending data to the server:", err)
			break
		}
	}
}

func (c *Client) read() {
	reader := bufio.NewReader(c.conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from the server:", err)
			os.Exit(1)
		}

		fmt.Print(line)
	}
}
