package main

import (
	"github.com/thanhtuan260593/todo/server"
)

func main() {
	server := server.New()
	server.Start()
}
