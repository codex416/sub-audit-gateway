package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/codex416/sub-audit-gateway/database"
)



func SaveAuditMiddleware() gin.HandlerFunc {


	return func(c *gin.Context) {


		start := time.Now()


		c.Next()



		token := c.Param("token")


		log := database.AuditLog{

			Token: token,

			IP: GetClientIP(c),

			UserAgent: c.GetHeader("User-Agent"),

			Path: c.Request.URL.Path,

			Status: c.Writer.Status(),

			CreatedAt: start,

		}



		if database.DB != nil {

			database.DB.Create(&log)

		}


	}

}
