// +build ignore

package main

import (
	"github.com/SunRunAway/go_httpauth"
	"net/http"
)

var s = go_httpauth.NewDigest("MyRealm", func(user, realm string) string {
	// Replace this with a real lookup function here
	return go_httpauth.CalculateHA1(user, realm, "mypass")
})

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if !s.Auth(w, r) {
		return
	}
	w.Write([]byte("<html><title>Hello</title><h1>Hello</h1></html>"))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
