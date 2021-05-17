package GCodeController

import "net"

const (
	GRBL = iota
)

type GCodeController struct {
	Type           int
	AvailablePorts []string
	Port           string
	Baudrate       string
	io             net.Conn
}

type RemoteController interface {
	Connect() (bool, error)
	Disconnect() (bool, error)
	Command() error
	ListPorts() error
	Write() error
	WriteLn() error
}
