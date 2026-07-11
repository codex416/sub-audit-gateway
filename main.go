package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/codex416/sub-audit-gateway/config"
	"github.com/codex416/sub-audit-gateway/database"
	"github.com/codex416/sub-audit-gateway/handler"
	"github.com/codex416/sub-audit-gateway/middleware"
)


func main() {


	cfg, err := config.LoadConfig("config/config.yaml")


	if err != nil {

		log.Fatal("Config load failed:", err)

	}



	err = database.InitDatabase()


	if err != nil {

		log.Fatal("Database init failed:", err)

	}



	r := gin.Default()



	// 请求审计

	r.Use(
		middleware.AuditMiddleware(),
	)



	// 保存数据库

	r.Use(
		middleware.SaveAuditMiddleware(),
	)



	r.GET("/", func(c *gin.Context) {


		c.JSON(200, gin.H{

			"status": "ok",

			"name": "sub-audit-gateway",

			"version": "v1.0",

		})


	})



	r.GET(
		"/sub/:token",
		handler.SubscribeHandler(cfg),
	)



	port := fmt.Sprintf(
		":%d",
		cfg.Server.Port,
	)



	fmt.Println(
		"Sub Audit Gateway running",
		port,
	)



	r.Run(port)

}
