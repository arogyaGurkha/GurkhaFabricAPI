package main

import (
	"github.com/arogyaGurkha/GurkhaFabricAPI/docs"
	lc "github.com/arogyaGurkha/GurkhaFabricAPI/lifecycle"
	"github.com/arogyaGurkha/GurkhaFabricAPI/peer"
	repo "github.com/arogyaGurkha/GurkhaFabricAPI/repository"
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

	// peer routes
	router.GET("/fabric/peer/", peer.GetPeerVersion)

	// lifecycle routes
	router.GET("/fabric/lifecycle/admin", lc.GetAdmin)
	router.POST("/fabric/lifecycle/admin/:organization", lc.SetAdmin)
	router.POST("/fabric/lifecycle/package", lc.PackageCC)
	router.POST("/fabric/lifecycle/install/:package_name", lc.InstallCC)
	router.POST("fabric/lifecycle/approve", lc.ApproveCC)
	router.POST("fabric/lifecycle/commit", lc.CommitCC)
	router.GET("fabric/lifecycle/install", lc.QueryInstalledCC)
	router.GET("fabric/lifecycle/approve", lc.QueryApprovedCC)
	router.GET("fabric/lifecycle/commit/organizations", lc.QueryCommitReadiness)
	router.GET("fabric/lifecycle/commit", lc.QueryCommittedCC)

	// repository routes
	router.POST("fabric/repository/clone", repo.CloneCC)

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // Points to the API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(":8080")
}

func updateSwagger() {
	cmd := exec.Command("swag", "init", "-g", "main.go", "--output", "docs")
	cmd.Run()
}
