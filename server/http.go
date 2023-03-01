package server

import "github.com/gin-gonic/gin"

func NewHttp() (server *gin.Engine) {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	return r
}
