package peplink

import "time"

type options struct {
	httpBasicEndpoint string
	httpClientSecret  string
	httpClientID      string
	timeout           time.Duration
	snmpAddress       string
	snmpCommunity     string
}
type Option func(*options) error

func WithHTTPBasicURL(url string) Option {
	return func(o *options) error {
		o.httpBasicEndpoint = url
		return nil
	}
}

func WithHTTPBasicClientID(id string) Option {
	return func(o *options) error {
		o.httpClientID = id
		return nil
	}
}

func WithHTTPBasicClientSecret(secret string) Option {
	return func(o *options) error {
		o.httpClientSecret = secret
		return nil
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *options) error {
		o.timeout = timeout
		return nil
	}
}

func WithSNMPAddress(address string) Option {
	return func(o *options) error {
		o.snmpAddress = address
		return nil
	}
}

func WithSNMPCommunity(community string) Option {
	return func(o *options) error {
		o.snmpCommunity = community
		return nil
	}
}
