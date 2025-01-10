package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func goTitle() {
	var (
		address                = "10.100.101.181:9923"
		protocol               = "tcp"
		pushGFX3               = "RGMOS:2\n"
		pullGFX3               = "RGMOH:2\n"
		duration time.Duration = 3
	)
	// Установите соединение с Livestream Studio
	conn, err := net.Dial(protocol, address)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(pushGFX3))
	if err != nil {
		fmt.Println("Error sending command:", err)
		return
	}

	time.Sleep(time.Second * duration)
	conn.SetReadDeadline(time.Now().Add(time.Second))

	_, err = conn.Write([]byte(pullGFX3))
	if err != nil {
		fmt.Println("Error sending push command:", err)
		return
	}

	// // Read the response from the server
	reader := bufio.NewReader(conn)
	reader.ReadString('\n')
}
