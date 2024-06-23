package server

import (
	"net"
)

type Router struct {
	handler map[string]handler
}

func NewRouter() *Router {
	return &Router{
		handler: make(map[string]handler),
	}
}

func (r *Router) HandlerFunc(cmd string, handler handler) {
	r.handler[cmd] = handler
}

func (r *Router) Listen(address string) error {
	server, err := NewServer(r, address)

	server.Listen()

	return err
}

func (r *Router) ExecuteCmd(conn net.Conn, cmd, args string) error {
	if fn, ok := r.handler[cmd]; ok {
		fn(conn, args)
	} else {
		_, err := conn.Write([]byte(ErrorReq))
		if err != nil {
			return err
		}
	}

	return nil
}
