package main

import (
	"fmt"
	"net/http"

	"github.com/kataras/realip"
	"github.com/kataras/tunnel"
)

func main() {
	opts := &realip.Options{
		Headers: []string{"X-Forwarded-For"},
		// PrivateSubnets: []realip.Range{
		// 	{
		// 		Start: net.ParseIP("192.168.0.0"),
		// 		End:   net.ParseIP("192.168.255.255"),
		// 	},
		// },
		// OR use `AddRange` instead:
	}
	opts.AddRange("192.168.0.0", "192.168.255.255")
	http.HandleFunc("/", handler(opts))

	srv := &http.Server{Addr: ":8080"}
	go fmt.Printf("• Public Address: %s\n", tunnel.MustStart(tunnel.WithServers(srv)))
	srv.ListenAndServe()
}

func handler(opts *realip.Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := opts.Get(r)
		fmt.Fprintf(w, "Your Public IPv4 is: %s", ip)
	}
}
