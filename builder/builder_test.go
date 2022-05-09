package builder

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPRequest_Build(t *testing.T) {
	request, err := NewHTTPRequest("https://example.com").
		AddHeader("User-Agent", "Golang patterns").
		Build()
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com", request.URL.String())
	assert.Equal(t, "Golang patterns", request.Header.Get("User-Agent"))
	assert.Equal(t, http.MethodGet, request.Method)
}
