package peer

import (
	"github.com/arogyaGurkha/GurkhaFabricAPI"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPeerVersion(t *testing.T) {
	router := main.setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/fabric/peer/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
