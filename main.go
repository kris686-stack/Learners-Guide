// servmux.go2 project main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "HELLO")
	})
	mux.HandleFunc("/hello/go", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "GO DHONI")
	})
	http.ListenAndServe(":9000", mux)
}
