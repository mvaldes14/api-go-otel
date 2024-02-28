package api

import "net/http"

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
