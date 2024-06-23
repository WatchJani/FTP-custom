package server

import (
	"log"
	"net"
)

type handler func(net.Conn, string)

const ErrorReq = "Wrong command\n"

type Server struct {
	addr           string
	listener       net.Listener
	readBufferSize int
	*Router
}

func NewServer(router *Router, address string) (*Server, error) {
	ls, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Server{
		listener:       ls,
		addr:           address,
		readBufferSize: 4096,
		Router:         router,
	}, nil
}

func (s *Server) Listen() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
		}

		s.addr = conn.RemoteAddr().String()

		go s.ReadLoop(conn)
	}
}

func (s *Server) ReadLoop(conn net.Conn) {
	defer conn.Close()
	payload := make([]byte, s.readBufferSize)

	for {
		n, err := conn.Read(payload)
		if err != nil {
			log.Println(err)
			return
		}

		cmd, args := Parser(payload[:n-1]) //Remove \n

		if fn, ok := s.handler[cmd]; ok {
			fn(conn, args)
		} else {
			_, err := conn.Write([]byte(ErrorReq))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
