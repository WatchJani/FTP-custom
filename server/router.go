package server

import (
	"net"
	"root/server/trie"
)

type Router struct {
	direct map[string]handler //direct rout
	trie.Trie[handler]
}

func NewRouter() *Router {
	return &Router{
		direct: make(map[string]handler),
		Trie:   trie.NewTrie[handler](),
	}
}

// test this func
func AnalyzeEndpoint(endpoint string) ([]string, bool) {
	pointer, res, dynamic := 0, make([]string, 0), false

	for i := 0; i < len(endpoint); i++ {
		if endpoint[i] == '/' {
			res = append(res, endpoint[pointer:i])
			pointer = i
		}

		if endpoint[i] == '{' {
			dynamic = !dynamic
		}
	}

	if pointer != len(endpoint) {
		res = append(res, endpoint[pointer:])
	}

	return res, dynamic
}

func (r *Router) HandlerFunc(cmd string, handler handler) {
	// path, isDynamic := AnalyzeEndpoint(cmd)
	// if isDynamic {
	// 	r.Insert(path, handler)
	// 	return
	// }

	r.direct[cmd] = handler
}

func (r *Router) Listen(address string) error {
	server, err := NewServer(r, address)

	server.Listen()

	return err
}

func (r *Router) ExecuteCmd(conn net.Conn, cmd, args string) error {
	if fn, ok := r.direct[cmd]; ok {
		fn(conn, args)
	} else {
		_, err := conn.Write([]byte(ErrorReq))
		if err != nil {
			return err
		}
	}

	return nil
}
