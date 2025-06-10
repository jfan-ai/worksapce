package main

import (
	"worksapce/webook/internal/web"
)

func main() {
	server := web.RegisterServer()
	server.Run(":8080")
}
