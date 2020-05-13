// mockHttpServer project main.go
package main

import (
	"fmt"
	"io"
	"net/http"
)

type display struct{}

func (h display) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("hello\n")
	res.Write(data)
	sasi := []byte("unni malu\n")
	res.Write(sasi)
	io.WriteString(res, "IN CANADA\n")
	fmt.Fprint(res, "Welcome to canada ,nice to meet you \n")
}

func main() {
	handler := display{}
	http.ListenAndServe(":8000", handler)
}
