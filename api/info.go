package handler

import (
	"fmt"
	"net/http"
)

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Section for Handler2</h1>")
}
