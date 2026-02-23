package controller

import (
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	var a []string = make([]string, 1, 6)
	fmt.Println(a)
	w.Write([]byte("test success"))
}
