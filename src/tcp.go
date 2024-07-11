package main

import (
	"fmt"
	"net"
	"time"
	"flag"
)

func main() {

	var ip string
	var port int
	var repeat int
	var wait int

	flag.StringVar(&ip, "ip", "", "IP address to connect to")
	flag.IntVar(&port, "port", 445, "Port number to connect to")
	flag.IntVar(&repeat, "repeat", 10, "Repeat  of connection")
	flag.IntVar(&wait, "wait", 1, "Waiting time for close connection")


	flag.Parse()

	if ip == "" {
		ip = "127.0.0.1"
	}

	address := fmt.Sprintf("%s:%d", ip, port)

	var connections []net.Conn


	for i := 1; i <= repeat; i++ {
		fmt.Printf("Connecting to %s, attempt %d/%d\n", address, i, repeat)

		conn, err := net.DialTimeout("tcp", address, 10*time.Second)
		if err != nil {
			fmt.Printf("Connection error: %v\n", err)
			continue
		}

		connections = append(connections, conn)
	}

	// 等待一分钟
	fmt.Printf("Wait %d minutes and close all connections. \n", wait)
	time.Sleep(time.Duration(wait) * time.Minute)

	// 关闭所有连接
	for _, conn := range connections {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection: %v\n", err)
		}
	}

	fmt.Println("All connections closed.")
}
