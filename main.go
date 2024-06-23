package main

import (
	"log"
	"net"
	s "root/server"
	"time"
)

func main() {

	r := s.NewRouter()

	r.HandlerFunc("QUIT", Quit)

	if err := r.Listen("localhost:5000"); err != nil {
		log.Println(err)
	}
}

func Quit(c net.Conn, args string) {
	defer c.Close()

	c.Write([]byte("log out: " + time.Now().String()))
}
