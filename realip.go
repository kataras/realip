package realip

import (
	"net"
	"net/http"
	"strings"
)

// Options holds the request `Headers` which IP should be fetched from.
// A header value should be separated by comma if contains more than one ip address.
// The `PrivateSubnets` field can be used to skip "local" addresses parsed by the `Headers` field.
//
// See `AddHeader`, `AddRange` and `Get` methods.
type Options struct {
	Headers        []string `json:"headers" yaml:"Headers" toml:"Headers"`
	PrivateSubnets []Range  `json:"privateSubnets" yaml:"PrivateSubnets" toml:"PrivateSubnets"`
}

// Default is an `Options` value with some default headers and private subnets.
// See `Get` method.
var Default = Options{
	Headers: []string{
		"X-Real-Ip",
		"X-Forwarded-For",
		"CF-Connecting-IP",
	},
	PrivateSubnets: []Range{
		{
			Start: net.ParseIP("10.0.0.0"),
			End:   net.ParseIP("10.255.255.255"),
		},
		{
			Start: net.ParseIP("100.64.0.0"),
			End:   net.ParseIP("100.127.255.255"),
		},
		{
			Start: net.ParseIP("172.16.0.0"),
			End:   net.ParseIP("172.31.255.255"),
		},
		{
			Start: net.ParseIP("192.0.0.0"),
			End:   net.ParseIP("192.0.0.255"),
		},
		{
			Start: net.ParseIP("192.168.0.0"),
			End:   net.ParseIP("192.168.255.255"),
		},
		{
			Start: net.ParseIP("198.18.0.0"),
			End:   net.ParseIP("198.19.255.255"),
		},
	},
}

// AddRange adds a private subnet to "opts".
// Should be called before any use of `Get`.
func (opts *Options) AddRange(start, end string) *Options {
	opts.PrivateSubnets = append(opts.PrivateSubnets, Range{
		Start: net.ParseIP(start),
		End:   net.ParseIP(end),
	})
	return opts
}

// AddHeader adds a proxy remote address header to "opts".
// Should be called before any use of `Get`.
func (opts *Options) AddHeader(headerKey string) *Options {
	opts.Headers = append(opts.Headers, headerKey)
	return opts
}

// Get extracts the real client's remote IP Address.
//
// Based on proxy headers of `Headers` and `PrivateSubnets`.
//
// Fallbacks to the request's `RemoteAddr` field which is filled by the server.
func (opts *Options) Get(r *http.Request) string {
	for _, headerKey := range opts.Headers {
		ipAddresses := strings.Split(r.Header.Get(headerKey), ",")
		if ip, ok := GetIPAddress(ipAddresses, opts.PrivateSubnets); ok {
			return ip
		}
	}

	addr := strings.TrimSpace(r.RemoteAddr)
	if addr != "" {
		if ip, _, err := net.SplitHostPort(addr); err == nil {
			return ip
		}
	}

	return addr
}

// Get is a shortcut of `Default.Get`.
// Extracts the real client's remote IP Address.
func Get(r *http.Request) string {
	return Default.Get(r)
}
