# Real IP

[![build status](https://img.shields.io/travis/com/kataras/realip/master.svg?style=for-the-badge&logo=travis)](https://travis-ci.com/github/kataras/realip) [![report card](https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/kataras/realip) [![godocs](https://img.shields.io/badge/go-%20docs-488AC7.svg?style=for-the-badge)](https://godoc.org/github.com/kataras/realip)

Extract the real HTTP client's Remote IP Address.

## Installation

The only requirement is the [Go Programming Language](https://golang.org/dl).

```sh
$ go get github.com/kataras/realip
```

## Getting Started

The main function is `Get`, it makes use of the `Default` options to extract the request's remote address.

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/kataras/realip"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    ip := realip.Get(r)
    fmt.Fprintf(w, "Your Public IPv4 is: %s", ip)
}
```

> The `Get(r)` functions calls the `Default.Get(r)` method

### Options

Here are the default values:

```go
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
```

Use the `AddRange` method helper to add an IP range in **custom** options:

```go
func main() {
    myOptions := &realip.Options{Headers: []string{"X-Forwarded-For"}}
    myOptions.AddRange("192.168.0.0", "192.168.255.255")

    // [...]
    http.HandleFunc("/", handler(opts))
}

func handler(opts *realip.Options) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request){
        ip := opts.Get(r)

        // [...]
    }
}
```

Please navigate through [_examples](_examples) directory for more.

## License

This software is licensed under the [MIT License](LICENSE).
