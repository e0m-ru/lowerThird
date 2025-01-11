package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	address  = flag.String("address", "10.100.101.181:9923", "server address")
	protocol = flag.String("protocol", "tcp", "network protocol")
	pushGFX3 = flag.String("pushGFX3", "RGMOS:2\n", "push command")
	pullGFX3 = flag.String("pullGFX3", "RGMOH:2\n", "pull command")
	duration = flag.Duration("duration", 3*time.Second, "duration to wait before pulling")
)

func goTitle() {
	// Установите соединение с Livestream Studio
	conn, err := net.Dial(*protocol, *address)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(*pushGFX3))
	if err != nil {
		fmt.Println("Error sending command:", err)
		return
	}

	time.Sleep(time.Second * *duration)
	conn.SetReadDeadline(time.Now().Add(time.Second))

	_, err = conn.Write([]byte(*pullGFX3))
	if err != nil {
		fmt.Println("Error sending push command:", err)
		return
	}

	// // Read the response from the server
	reader := bufio.NewReader(conn)
	reader.ReadString('\n')
}
