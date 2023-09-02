package peplink

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type apiEnvelope struct {
	Stat     string          `json:"stat"`              // API call status {ok, fail}
	Response json.RawMessage `json:"response"`          // Any additional information of the success call will be here
	Code     int             `json:"code,omitempty"`    // Error code of the API call (only appears if the API call is not successful)
	Message  string          `json:"message,omitempty"` // Error message of the API call (only appears if the API call is not successful)
	Notice   interface{}     `json:"notice,omitempty"`  // Extra information about the API request (not part of the normal response)
}

func (c *Client) doRequest(ctx context.Context, endpoint, method string, body any) (json.RawMessage, error) {
	envelope := &apiEnvelope{}

	request := c.httpClient.R().
		SetContext(ctx).
		SetResult(envelope).
		SetError(envelope)

	var (
		err error
	)
	switch method {
	case http.MethodGet:
		_, err = request.Get(endpoint)
	case http.MethodPost:
		_, err = request.
			SetBody(body).
			Post(endpoint)
	default:
		return nil, fmt.Errorf("unknown method: %s Only GET and POST supported", method)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to do HTTP request: %w", err)
	}

	if envelope.Stat != "ok" {
		return nil, fmt.Errorf("status of the request isn't ok: stat='%s' body='%v'", envelope.Stat, envelope)
	}

	return envelope.Response, nil
}
