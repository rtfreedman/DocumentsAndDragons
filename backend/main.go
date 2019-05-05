package main

import (
	"flag"

	"github.com/rtfreedman/DocumentsAndDragons/backend/api"
	_ "github.com/rtfreedman/DocumentsAndDragons/backend/store"
)

var port int
var external bool

func main() {
	flag.IntVar(&port, "port", 8010, "Port to run API on")
	flag.BoolVar(&external, "external", false, "Run as externally accessible API")
	flag.Parse()
	api.RunAPI(port, external)
}
