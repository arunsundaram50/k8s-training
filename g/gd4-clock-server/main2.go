package main

import (
	"io"
	"log"
	"net"
	"time"
)

func writeTime(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format(time.UnixDate)+"\n")
		if err != nil {
			log.Panic(err)
		}
		time.Sleep(time.Second)
	}
}

// One-client listener (for the life of the server)
func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Panic(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Panic(err)
	}

	writeTime(conn)
}
