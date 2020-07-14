package main

import (
	"fmt"
	"net/http"

	"github.com/kataras/realip"
)

func main() {
	http.HandleFunc("/", defaultOptions)
	http.ListenAndServe(":8080", nil)
}

func defaultOptions(w http.ResponseWriter, r *http.Request) {
	ip := realip.Get(r)
	fmt.Fprintf(w, "Your Public IPv4 is: %s", ip)
}
