package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/get", Get)
	http.ListenAndServe(":9090", nil)
}
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "wowww Getrpoduct api is called")
}
