package rest

import "net"

// Option -.
type Option func(*RestServer)

// Port -.
func Port(port string) Option {
	return func(s *RestServer) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}
