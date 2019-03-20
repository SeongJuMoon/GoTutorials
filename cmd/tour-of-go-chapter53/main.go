package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punch    string
	Who      string
}

func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, str)
}

func (structs *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, structs.Greeting+structs.Punch+structs.Who)
}

func main() {

	http.Handle("/string", String("I ' m seongju moon"))
	http.Handle("/struct", &Struct{"Hello", ":", "Gopher"})

	http.ListenAndServe("localhost:4000", nil)
}
