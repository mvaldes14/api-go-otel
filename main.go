package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello World"))
}

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("/",indexHandler)


  http.ListenAndServe(":8080", mux)

}
