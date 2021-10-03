package server

type Server interface {
	start() error
}

type ServerConfig struct {
}
