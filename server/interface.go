package server

type Server interface {
	Start() error
}

type ServerConfig struct {
}
