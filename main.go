package main

import (
	"github.com/arogyaGurkha/GurkhaFabricAPI/admin"
	cc "github.com/arogyaGurkha/GurkhaFabricAPI/chaincode"
	"github.com/arogyaGurkha/GurkhaFabricAPI/docs"
	lc "github.com/arogyaGurkha/GurkhaFabricAPI/lifecycle"
	"github.com/arogyaGurkha/GurkhaFabricAPI/network"
	"github.com/arogyaGurkha/GurkhaFabricAPI/peer"
	repo "github.com/arogyaGurkha/GurkhaFabricAPI/repository"
	"github.com/arogyaGurkha/GurkhaFabricAPI/repository/dashboard"
	search "github.com/arogyaGurkha/GurkhaFabricAPI/repository/search"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os/exec"
)

// @title Hyperledger Fabric Gurkhaman API
// @description API to run fabric binaries
// @version 0.1
// @contact.name arogya.Gurkha
// @contact.url https://github.com/arogyaGurkha
func main() {
	log.Println("==========Application Start==========")

	admin.SetConnection()

	updateSwagger()
	router := setupRouter()
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"content-type"},
	}))
	//jeho admin
	adminRouter := router.Group("/fabric/admin")
	{
		adminRouter.GET("/addorg", admin.AddOrg)
	}

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
	router.GET("fabric/repository/clone", repo.CloneCC)
	router.GET("fabric/repository/pull", repo.PullOrigin)
	router.POST("fabric/repository/revert", repo.RevertUpdate)
	router.POST("fabric/repository/reset", repo.ResetLocal)
	router.GET("fabric/repository/updates", repo.CheckUpdate)
	router.GET("fabric/repository/logs", repo.GetRefLogs)

	// chaincode routes
	router.POST("fabric/chaincode/invoke", cc.InvokeCC)
	router.GET("fabric/chaincode/query", cc.QueryCC)

	// network routes
	router.POST("fabric/network/up", network.StartFabricWChannel)
	router.POST("fabric/network/down", network.StopFabric)

	// repository/elasticsearch routes
	router.GET("fabric/dashboard/smart-contracts", search.ESSearchAll)
	router.GET("fabric/dashboard/smart-contracts/:id", search.EsDocumentByID)
	router.GET("fabric/dashboard/search", search.ESSearchWithLanguage)

	// repository/dashboard routes
	router.POST("fabric/dashboard/deployCC", dashboard.InstallWithDeployCC)
	router.POST("fabric/dashboard/smart-contracts", dashboard.AddDataToES)
	//router.POST("fabric/dashboard/smart-contracts/transaction", dashboard.CreateTransaction)
	router.POST("fabric/dashboard/smart-contracts/transaction", dashboard.AssetTransfer2)
	router.GET("fabric/dashboard/smart-contracts/asset", dashboard.QueryAssets)

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // Points to the API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func updateSwagger() {
	cmd := exec.Command("swag", "init", "-g", "main.go", "--output", "docs")
	cmd.Run()
}
