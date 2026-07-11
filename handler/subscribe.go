package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/codex416/sub-audit-gateway/config"
)


func SubscribeHandler(cfg *config.Config) gin.HandlerFunc {


	return func(c *gin.Context) {


		token := c.Param("token")


		if token == "" {

			c.JSON(400, gin.H{

				"error": "missing token",

			})

			return

		}



		upstream := strings.TrimRight(
			cfg.Subscription.Upstream,
			"/",
		)



		url := upstream + "/" + token



		resp, err := http.Get(url)


		if err != nil {

			c.JSON(502, gin.H{

				"error": "upstream unavailable",

			})

			return

		}


		defer resp.Body.Close()



		body, err := io.ReadAll(resp.Body)


		if err != nil {

			c.JSON(500, gin.H{

				"error": "read failed",

			})

			return

		}



		c.Data(

			resp.StatusCode,

			"text/plain; charset=utf-8",

			body,

		)


	}

}
