package main

import (
	"github.com/go-fuego/fuego"
  "fmt"
)

func newServer() *fuego.Server {
	s := fuego.NewServer()
	return s
}


func main() {
	server := newServer()

	fuego.Get(server, "/", func(c fuego.ContextNoBody) (fuego.HTML, error) {
		return c.Render("templates/index.html", "Hello, World!")
  })
	fuego.Get(server, "/api", func(c fuego.ContextNoBody) (string, error) {
    return fmt.Sprintf("{msg:hello}"), nil
  })
	
  server.Run()
}
