package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/heckphy/wb-l2/develop/dev10/telnet"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут на подключение (по умолчанию 10s)")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go run task.go <host> <port>")
		os.Exit(1)
	}

	host, port := args[0], args[1]

	client, err := telnet.NewClient(host, port, *timeout)
	if err != nil {
		fmt.Println("Error creating Telnet client:", err)
		os.Exit(1)
	}

	client.Start()
}
