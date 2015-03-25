package replicator

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("BAR"))
})

func TestEnforced(t *testing.T) {

	s := NewReplicator("http://localhost.com", false, 100) // make sure the request goes through, 100% of requests

	res := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "http://examle.com/foo", nil)

	s.ServeHTTP(res, req, testHandler)

	assert.Equal(t, res.Code, 200)
}

func TestPassive(t *testing.T) {
    s := NewReplicator("127.0.0.1", true, 10)

    res := httptest.NewRecorder()

    req, _ := http.NewRequest("POST", "http://example.com", nil)

    s.ServeHTTP(res, req, testHandler)

}
