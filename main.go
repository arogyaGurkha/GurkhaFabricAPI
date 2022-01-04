package main

import (
	"github.com/arogyaGurkha/GurkhaFabricAPI/docs"
	lc "github.com/arogyaGurkha/GurkhaFabricAPI/lifecycle"
	"github.com/arogyaGurkha/GurkhaFabricAPI/peer"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"os/exec"
)

// @title Hyperledger Fabric Gurkhaman API
// @description API to run fabric binaries
// @version 0.1
// @contact.name arogya.Gurkha
// @contact.url https://github.com/arogyaGurkha
func main() {
	updateSwagger()

	router := gin.Default()

	// Routes
	router.GET("/fabric/peer/", peer.GetPeerVersion)
	router.POST("/fabric/lifecycle/admin/:organization", lc.SetAdmin)
	router.GET("/fabric/lifecycle/admin", lc.GetAdmin)

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // Points to the API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("localhost:8080")
}

func updateSwagger() {
	cmd := exec.Command("swag", "init", "-g", "main.go", "--output", "docs")
	cmd.Run()
}
