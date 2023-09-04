package peplink

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

// Client for the https://www.peplink.com/ic2-api-doc
type Client struct {
	httpClient *resty.Client
	log        *slog.Logger
}

// NewClient creates a new Peplink Client and authenticates against the API
// Runs token update process in the background
func NewClient(ctx context.Context, opts ...Option) (*Client, error) {
	options := &options{
		timeout:           10 * time.Second,
		httpBasicEndpoint: "http://127.0.0.1:8080",
		snmpAddress:       "127.0.0.1:161",
		snmpCommunity:     "public",
	}

	for _, o := range opts {
		o(options)
	}

	rest := resty.New().
		SetBaseURL(options.httpBasicEndpoint).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetTimeout(options.timeout)

	c := &Client{
		httpClient: rest,
		log:        slog.Default(),
	}

	ttl, err := c.authenticate(context.Background(), options.httpClientID, options.httpClientSecret)
	if err != nil {
		c.log.Error("Failed to authenticate", "error", err)

		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}

	go func() {
		time.Sleep(ttl - 10*time.Minute)
		err := c.watchToken(ctx, options.httpClientID, options.httpClientSecret)
		if err != nil {
			c.log.Error("Failed to watch token", "error", err)
		}
	}()

	return c, nil
}

func (c *Client) watchToken(ctx context.Context, clientID, clientSecret string) error {
	c.log.Info("Peplink token refresh goroutine started")
	defer c.log.Info("Peplink token refresh goroutine stopped")

	ttl, err := c.authenticate(ctx, clientID, clientSecret)
	if err != nil {
		return fmt.Errorf("failed to update token: %w", err)
	}
	if ttl < 10*time.Minute {
		return fmt.Errorf("token TTL is too short: %s", fmt.Sprint(ttl))
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(ttl - 10*time.Minute):
			ttl, err = c.authenticate(ctx, clientID, clientSecret)
			if err != nil {
				return fmt.Errorf("failed to update token: %w", err)
			}
		}
	}
}

func (c *Client) authenticate(ctx context.Context, clientID, clientSecret string) (time.Duration, error) {
	type tokenRequest struct {
		ClientID     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
		Scope        string `json:"scope"`
	}
	type tokenResponse struct {
		Stat     string `json:"stat"`
		Response struct {
			// The access token string for API call
			AccessToken string `json:"accessToken"`
			// Expiration time in seconds
			ExpiresIn string `json:"expiresIn"`
		} `json:"response"`
	}
	resp := &tokenResponse{}

	rr, err := c.httpClient.NewRequest().SetBody(tokenRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        "api",
	}).
		SetResult(resp).
		SetContext(ctx).
		Post("/api/auth.token.grant")

	if err != nil {
		return 0, fmt.Errorf("failed to authenticate: %w", err)
	}

	if resp.Stat != "ok" {
		return 0, fmt.Errorf("failed to authenticate: stat='%s' body='%s'", resp.Stat, rr.Body())
	}

	ttlInt, err := strconv.Atoi(resp.Response.ExpiresIn)
	if err != nil {
		return 0, fmt.Errorf("unexpeted 'ExpiresIn': %w", err)
	}

	ttl := time.Duration(ttlInt) * time.Second

	c.log.Info("Authenticated against Peplink API", "status", rr.Status(), "TTL", fmt.Sprint(ttl))

	c.httpClient.SetPathParam("accessToken", resp.Response.AccessToken)

	return ttl, nil
}
