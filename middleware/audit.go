package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)


func AuditMiddleware() gin.HandlerFunc {


	return func(c *gin.Context) {


		start := time.Now()



		ip := c.ClientIP()


		userAgent := c.GetHeader("User-Agent")


		path := c.Request.URL.Path



		c.Next()



		duration := time.Since(start)



		status := c.Writer.Status()



		log.Printf(
			"[AUDIT] IP=%s PATH=%s UA=%s STATUS=%d TIME=%s",
			ip,
			path,
			userAgent,
			status,
			duration,
		)


	}

}
