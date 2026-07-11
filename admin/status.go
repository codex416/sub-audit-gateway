package admin

import (
	"time"

	"github.com/gin-gonic/gin"
)


func StatusHandler() gin.HandlerFunc {


	return func(c *gin.Context) {


		c.JSON(200, gin.H{

			"service": "sub-audit-gateway",

			"status": "running",

			"version": "v1.0",

			"time": time.Now().Format(
				"2006-01-02 15:04:05",
			),

		})


	}

}
