package peplink

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_authenticate(t *testing.T) {
	tests := []struct {
		name         string
		clientID     string
		clientSecret string
		response     string
		wantToken    bool
		wantErr      bool
	}{
		{"happy",
			"client_id",
			"client_secret",
			`{
				"stat": "ok",
				"response": {
					"accessToken": "43c65216eb16d779092fc40b184a1794",
					"authorizationType": "3",
					"scope": "api",
					"expiresIn": "172800"
				}
			}`,
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					require.Equal(t, "/api/auth.token.grant", r.URL.Path)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(tt.response))
				}),
			)

			defer srv.Close()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c, err := NewClient(ctx, Options{
				URL:          srv.URL,
				Timeout:      5 * time.Second,
				ClientID:     tt.clientID,
				ClientSecret: tt.clientSecret,
			})
			require.Equal(t, tt.wantErr, err != nil, "authenticate() error = %v, wantErr %v", err, tt.wantErr)

			if tt.wantToken {
				require.NotEmpty(t, c.httpClient.PathParams["accessToken"], "authenticate() token is empty")
			}
		})
	}
}
