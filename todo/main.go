package main

import (
	"github.com/thanhtuan260593/todo/web/server"
)

func main() {
	server := server.New()
	server.Start()
}
