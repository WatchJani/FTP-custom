package server

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
