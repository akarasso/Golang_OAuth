package main

import (
	"log"
	"net/http"

	"github.com/akarasso/oauth_server/handler"
)

/*
**	Router to use?
**	https://github.com/savsgio/atreugo
 */
func main() {
	hndlr := http.NewServeMux()
	hndlr.HandleFunc("/token", handler.Token)
	err := http.ListenAndServe(":9000", hndlr)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
