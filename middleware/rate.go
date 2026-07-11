package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/codex416/sub-audit-gateway/config"
	"github.com/codex416/sub-audit-gateway/notify"
)


type visitor struct {

	count int

	last time.Time

}


var (

	visitors = make(map[string]*visitor)

	mu sync.Mutex

)



func RateLimitMiddleware(cfg *config.Config) gin.HandlerFunc {


	return func(c *gin.Context) {


		ip := GetClientIP(c)



		mu.Lock()

		defer mu.Unlock()



		now := time.Now()



		v, exists := visitors[ip]



		if !exists {


			visitors[ip] = &visitor{

				count:1,

				last:now,

			}


			c.Next()

			return

		}



		// 一分钟窗口

		if now.Sub(v.last) > time.Minute {


			v.count = 1

			v.last = now


			c.Next()

			return

		}



		v.count++



		if v.count > 60 {


	notify.SendAlert(
		cfg,
		"IP访问超限",
		"IP: "+ip+" 请求频率超过限制",
	)



	c.AbortWithStatusJSON(

		http.StatusTooManyRequests,

		gin.H{

			"error":"too many requests",

		},

	)


	return

}



		c.Next()


	}

}
