package web

import "github.com/gin-gonic/gin"

func RegisterServer() *gin.Engine {
	server := gin.Default()
	registerUsersRoutes(server)
	return server
}
