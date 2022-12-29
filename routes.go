package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes() {

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	v1 := router.Group("/v1")
	{
		authed := v1.Group("/:auth_token")
		{
			effects := authed.Group("/effects")
			{
				effects.GET("/effectsList", getEffectsList)
				effects.PUT("", putEffect)
			}
		}
	}
}
