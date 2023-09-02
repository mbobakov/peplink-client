package peplink

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmespath/go-jmespath"
)

type firmware struct {
	Version  string `json:"version"`
	Bootable bool   `json:"bootable"`
	InUse    bool   `json:"inUse"`
}

// Returns the firmware version of the device
func (c *Client) FirmwareVersion(ctx context.Context) (string, error) {
	msg, err := c.doRequest(ctx, "/api/info.frw.version", http.MethodGet, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get firmware version via http: %w", err)
	}
	var buf interface{}

	err = json.Unmarshal(msg, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	orderI, err := jmespath.Search("order", buf)
	if err != nil {
		return "", fmt.Errorf("failed to get firmware version from json: %w", err)
	}
	for _, oi := range orderI.([]interface{}) {
		i, ok := oi.(float64)
		if !ok {
			return "", fmt.Errorf("failed to get firmware version: order is not a float64")
		}
		versionI, err := jmespath.Search(fmt.Sprintf("\"%d\"", int(i)), buf)
		if err != nil {
			return "", fmt.Errorf("failed to get firmware version from json: %w", err)
		}
		// Do marshal/unmarshal to workaround interface{} hussle
		buf, err := json.Marshal(versionI)
		if err != nil {
			return "", fmt.Errorf("failed to get firmware version from json: %w", err)
		}
		version := firmware{}
		err = json.Unmarshal(buf, &version)
		if err != nil {
			return "", fmt.Errorf("failed to get firmware version from json: %w", err)
		}
		if version.InUse {
			return version.Version, nil
		}
	}

	return "", fmt.Errorf("failed to get firmware version: no firmware in use")

}
