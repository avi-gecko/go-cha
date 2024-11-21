package rest_test

import (
	"testing"

	"github.com/avi-gecko/go-cha/internal/api"
)

func TestCreateRest(t *testing.T) {
	rest_test, err := api.Create("rest", "127.0.0.1", "8080")

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", api_test)
}
