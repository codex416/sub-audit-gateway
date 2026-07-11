package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"status": "ok",
			"name":   "sub-audit-gateway",
			"version": "v1.0",
		})

	})

	fmt.Println("Sub Audit Gateway running on :8080")

	r.Run(":8080")
}
