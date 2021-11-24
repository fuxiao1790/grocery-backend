package server

type Server interface {
	Start() error
}

type Config struct {
	Port     int
	CertFile string
	KeyFile  string
}
