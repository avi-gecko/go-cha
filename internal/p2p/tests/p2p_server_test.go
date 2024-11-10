package p2p_server_test

import (
	"testing"

	"github.com/avi-gecko/go-cha/internal/p2p"
)

func TestCreateIPv4(t *testing.T) {
	p2p_test, err := p2p.Create("ipv4", "127.0.0.1", "8080")

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", p2p_test)
}

func TestCreateIPv4WithWrongType(t *testing.T) {
	p2p_test, err := p2p.Create("ipv", "127.0.0.1", "8080")

	if err == nil {
		t.Error(err)
	}

	t.Logf("Error: %v, instance: %v", err, p2p_test)
}

func TestCreateIPv4WithWrongIP(t *testing.T) {
	p2p_test, err := p2p.Create("ipv4", "127230.0.1", "8080")

	if err == nil {
		t.Error(err)
	}

	t.Logf("Error: %v, instance: %v", err, p2p_test)
}

func TestCreateIPv4WithWrongPort(t *testing.T) {
	p2p_test, err := p2p.Create("ipv4", "127.0.0.1", "80000")

	if err == nil {
		t.Error(err)
	}

	t.Logf("Error: %v, instance: %v", err, p2p_test)
}

func TestP2PListen(t *testing.T) {
	p2p_test, err := p2p.Create("ipv4", "127.0.0.1", "8081")

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", p2p_test)
	//TODO: add client for test
}
