package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testServer *server

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func setup() {
	s := NewGroceryServer(&Config{
		PortHttp:  8080,
		PortHttps: 443,
		CertFile:  "./cert.pem",
		KeyFile:   "./key.pem",
	}, nil)

	testServer = s.(*server)

	go func() { testServer.Start() }()
}

func Test_GetItems_Invalid_Route(t *testing.T) {
	req := httptest.NewRequest("GET", "/INVALID_ROUTE", nil)
	res, err := testServer.app.Test(req, 1)
	assert.Nil(t, err)
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
}
