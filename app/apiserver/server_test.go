package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIserver_indexHandler(t *testing.T) {
	server := New(":8080")
	recorder := httptest.NewRecorder()
	request,_ := http.NewRequest(http.MethodGet, "/test", nil)
	server.indexHandler().ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "its test")
}
