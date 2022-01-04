package lifecycle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type currentAdmin struct {
	Admin string `json:"admin"`
}

// SetAdmin
// @Summary Set an org as the admin.
// @Description Use terminal environmental variables to set the admin for peer cli container. Only Org1 and Org2 are supported.
// @Accept json
// @Param organization path string true "organization to be set as admin (Org1 and Org2 supported)"
// @Produce json
// @Tags lifecycle
// @Success 200 {object} currentAdmin
// @Router /fabric/lifecycle/admin/{organization} [post]
func SetAdmin(c *gin.Context) {

	os.Setenv("CORE_PEER_TLS_ENABLED", "true")

	var admin currentAdmin

	organization := c.Param("organization")

	if organization == "Org1" {
		os.Setenv("CORE_PEER_ADMIN", organization)
		os.Setenv("CORE_PEER_LOCALMSPID", "Org1MSP")
		os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE", "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt")
		os.Setenv("CORE_PEER_MSPCONFIGPATH", "${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp")
		os.Setenv("CORE_PEER_ADDRESS", "localhost:7051")

		admin.Admin = organization
		c.IndentedJSON(http.StatusOK, admin)
		return
	}

	if organization == "Org2" {
		os.Setenv("CORE_PEER_ADMIN", organization)
		os.Setenv("CORE_PEER_LOCALMSPID", "Org2MSP")
		os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE", "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt")
		os.Setenv("CORE_PEER_MSPCONFIGPATH", "${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp")
		os.Setenv("CORE_PEER_ADDRESS", "localhost:9051")

		admin.Admin = organization
		c.IndentedJSON(http.StatusOK, admin)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Setting admin failure."})
}

// GetAdmin
// @Summary Get the current admin org.
// @Description Use terminal environmental variables to get the admin for peer cli container. Only Org1 and Org2 are supported.
// @Accept json
// @Produce json
// @Tags lifecycle
// @Success 200 {object} currentAdmin
// @Router /fabric/lifecycle/admin [get]
func GetAdmin(c *gin.Context) {
	var admin currentAdmin

	envAdmin := os.Getenv("CORE_PEER_ADMIN")

	if envAdmin == "Org1" || envAdmin == "Org2" {
		admin.Admin = envAdmin
		c.IndentedJSON(http.StatusOK, admin)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error getting current admin."})

}
