package waf

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasConnectionUpgrade(t *testing.T) {
	t.Parallel()

	tests := []struct {
		connection string
		expected   bool
	}{
		{"", false},
		{"keep-alive", false},
		{"Upgrade", true},
		{"upgrade", true},
		{"UPGRADE", true},
		{"keep-alive, Upgrade", true},
		{" upgrade ", true},
		{"upgrades", false},
	}

	for _, tt := range tests {
		t.Run(tt.connection, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.connection != "" {
				req.Header.Set("Connection", tt.connection)
			}
			assert.Equal(t, tt.expected, HasConnectionUpgrade(req))
		})
	}
}
