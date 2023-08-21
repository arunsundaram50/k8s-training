package main

import (
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func writeTime(conn net.Conn, clientID int) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, strconv.Itoa(clientID)+": "+time.Now().Format(time.UnixDate)+"\n")
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}

// Multi-client clock-server
func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Panic(err)
	}

	clientID := 0
	for {
		clientID += 1
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go writeTime(conn, clientID)
	}
}
