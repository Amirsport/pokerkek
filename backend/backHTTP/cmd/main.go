package main

import "pokerok/internal/server"

func main() {
	srv := server.Server{}
	srv.StartServer()
}
