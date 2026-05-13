package main

import (
	gorrm "RatWAF/database"
	net "RatWAF/network"
	"RatWAF/web"
)

func main() {
	gorrm.Conn()
	go net.ReverseProxy()

	r := web.InitRouter()
	r.Run(":9090")
}
