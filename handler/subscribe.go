package handler

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/codex416/sub-audit-gateway/config"
)



func SubscribeHandler(cfg *config.Config) gin.HandlerFunc {


	return func(c *gin.Context) {


		token := c.Param("token")
		
if len(token) < 5 {

	c.JSON(
		400,
		gin.H{
			"error":"invalid token",
		},
	)

	return

}


		if token == "" {


			c.JSON(
				400,
				gin.H{
					"error":"missing token",
				},
			)


			return

		}



		upstream := strings.TrimRight(
			cfg.Subscription.Upstream,
			"/",
		)



		target := upstream + "/" + token



		req, err := http.NewRequest(
			"GET",
			target,
			nil,
		)


		if err != nil {

			c.JSON(
				500,
				gin.H{
					"error":"create request failed",
				},
			)

			return

		}



		// 转发客户端信息

		req.Header.Set(
			"User-Agent",
			c.GetHeader("User-Agent"),
		)



		client := &http.Client{

	Timeout: 10 * time.Second,

}


resp, err := client.Do(req)



		if err != nil {


			c.JSON(
				502,
				gin.H{
					"error":"upstream unavailable",
				},
			)


			return

		}



		defer resp.Body.Close()



		body, err := io.ReadAll(
			resp.Body,
		)



		if err != nil {


			c.JSON(
				500,
				gin.H{
					"error":"read response failed",
				},
			)


			return

		}



		// 保留订阅信息头

		if value := resp.Header.Get("subscription-userinfo"); value != "" {


			c.Header(
				"subscription-userinfo",
				value,
			)


		}



		c.Data(
			resp.StatusCode,
			"text/plain; charset=utf-8",
			body,
		)


	}

}
