package controllers

import (
	"github.com/gin-gonic/gin"
)

func ShowWisdom(c *gin.Context) {
	// Return the info response as JSON
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "The SiliconeOracle says...",
	})

}
