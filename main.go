package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/codex416/sub-audit-gateway/config"
	"github.com/codex416/sub-audit-gateway/handler"
)


func main() {


	cfg, err := config.LoadConfig("config/config.yaml")


	if err != nil {

		log.Fatal("Config load failed:", err)

	}



	r := gin.Default()



	// 健康检查

	r.GET("/", func(c *gin.Context) {


		c.JSON(200, gin.H{

			"status": "ok",

			"name": "sub-audit-gateway",

			"version": "v1.0",

		})


	})



	// 订阅代理入口

	r.GET("/sub/:token",
		handler.SubscribeHandler(cfg),
	)



	port := fmt.Sprintf(":%d", cfg.Server.Port)



	fmt.Println(
		"Sub Audit Gateway running",
		port,
	)



	r.Run(port)

}
