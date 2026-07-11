package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)



func AuditMiddleware() gin.HandlerFunc {


	return func(c *gin.Context) {


		start := time.Now()



		ip := GetClientIP(c)


		ipVersion := IPVersion(ip)


		userAgent := c.GetHeader("User-Agent")


		path := c.Request.URL.Path



		c.Next()



		status := c.Writer.Status()


		duration := time.Since(start)



		log.Printf(
			"[AUDIT] IP=%s VERSION=%s PATH=%s UA=%s STATUS=%d COST=%s",
			ip,
			ipVersion,
			path,
			userAgent,
			status,
			duration,
		)


	}

}
