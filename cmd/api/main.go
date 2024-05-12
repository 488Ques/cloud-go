package main

import "cloud-go/internal/server"

func main() {
	s := server.New()
	s.Run()
}
