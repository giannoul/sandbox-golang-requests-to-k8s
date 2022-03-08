package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	name := "localhost"
	addrs, err := net.LookupHost(name)
	checkError(err)

	addr := net.ParseIP(addrs[0])
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}

	service := "localhost:8080"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	tcpConnectionDeadlineError := conn.SetDeadline(time.Now().Add(time.Second * 180))
	checkError(tcpConnectionDeadlineError)

	tcpConnectionReadDeadlineError := conn.SetReadDeadline(time.Now().Add(time.Second * 180))
	checkError(tcpConnectionReadDeadlineError)

	tcpConnectionWriteDeadlineError := conn.SetWriteDeadline(time.Now().Add(time.Second * 2))
	checkError(tcpConnectionWriteDeadlineError)

	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: localhost:8080\r\n\r\n"))
	checkError(err)

	//result, err := readFully(conn)
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}
