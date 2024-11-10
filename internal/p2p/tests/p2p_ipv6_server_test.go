package p2p_test

import (
	"testing"

	"github.com/avi-gecko/go-cha/internal/p2p"
)

func TestCreateIPv6(t *testing.T) {
	p2p_test, err := p2p.Create("ipv6", "0:0:0:0:0:0:0:1", "8080")

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", p2p_test)
}

func TestCreateIPv6WithWrongIP(t *testing.T) {
	p2p_test, err := p2p.Create("ipv6", "0:0:0:0:0:0:0", "8080")

	if err == nil {
		t.Error(err)
	}

	t.Logf("Error: %v, instance: %v", err, p2p_test)
}

func TestCreateIPv6WithWrongPort(t *testing.T) {
	p2p_test, err := p2p.Create("ipv6", "0:0:0:0:0:0:0:1", "80000")

	if err == nil {
		t.Error(err)
	}

	t.Logf("Error: %v, instance: %v", err, p2p_test)
}

func TestP2PIPv6Listen(t *testing.T) {
	p2p_test, err := p2p.Create("ipv6", "0:0:0:0:0:0:0:1", "8081")

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", p2p_test)
	//TODO: add client for test
}
