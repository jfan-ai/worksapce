package main

import "worksapce/webook2/internal/web"

func main() {
	server := web.RegisterRoutes()
	server.Run(":8080")
}
