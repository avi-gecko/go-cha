package p2p

import (
	"errors"
	"net"
)

type P2P interface {
	Listen() error
	handleConnection(net.Conn) error
}

func Create(p2p_type string, ip string, port string) (P2P, error) {
	var p2p P2P

	switch p2p_type {
	case "ipv4":
		server, err := net.Listen("tcp", ip+":"+port)

		if err != nil {
			return nil, errors.New("Can' create p2p with type: " + p2p_type + "\nCause: " + err.Error())
		}

		p2p = &P2PIPv4{ipv4: ip, port: port, server: server}
	//TODO: add ipv6
	default:
		return nil, errors.New("Can' create p2p with type: " + p2p_type + "\nCause: incorrect type")
	}

	return p2p, nil
}

type P2PIPv4 struct {
	ipv4   string
	port   string
	server net.Listener
	client net.Conn
}

func (p2p *P2PIPv4) handleConnection(conn net.Conn) error {
	//TODO: add output to stdin or socket for gui
	return nil
}

func (p2p *P2PIPv4) Listen() error {
	for {
		conn, err := p2p.server.Accept()
		if err != nil {
			return errors.New("P2P server can't listen\nCause: " + err.Error())
		}
		go p2p.handleConnection(conn)
	}
}

type P2PIPv6 struct {
	ipv6   string
	port   string
	server net.Listener
	client net.Conn
}
