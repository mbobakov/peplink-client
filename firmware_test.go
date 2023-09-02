package peplink

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
)

func TestClient_FirmwareVersion(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     string
		wantErr  bool
	}{
		{"happy",
			`{
			"stat": "ok",
			"response": {
			  "1": {
				"version": "8.2.0s036 build 4979",
				"bootable": true,
				"inUse": false
			  },
			  "2": {
				"version": "8.3.0 build 5229",
				"bootable": true,
				"inUse": true
			  },
			  "order": [
				1,
				2
			  ]
			}
		  }`,
			"8.3.0 build 5229",
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					require.Equal(t, "/api/info.frw.version", r.URL.Path)
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(tt.response))
				}),
			)

			defer srv.Close()

			c := Client{
				httpClient: resty.New().
					SetBaseURL(srv.URL).
					SetHeader("Content-Type", "application/json").
					SetHeader("Accept", "application/json"),
				log: slog.Default(),
			}

			got, err := c.FirmwareVersion(context.Background())
			require.Equal(t, tt.wantErr, err != nil, "FirmwareVersion() error = %v, wantErr %v", err, tt.wantErr)
			require.Equal(t, tt.want, got)
		})
	}
}
