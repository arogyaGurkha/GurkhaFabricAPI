package lifecycle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

type approvedChaincodeResponse struct {
	PackageID         string `json:"package_ID"`
	Sequence          int32  `json:"sequence"`
	Version           string `json:"version"`
	InitRequired      bool   `json:"init_required"`
	EndorsementPlugin string `json:"endorsement_plugin"`
	ValidationPlugin  string `json:"validation_plugin"`
}

type installedChaincodeResponse struct {
	PackageID string `json:"package_ID"`
	Label     string `json:"label"`
}

type committedChaincodeResponse struct {
	Sequence          int32           `json:"sequence"`
	Version           string          `json:"version"`
	EndorsementPlugin string          `json:"endorsement_plugin"`
	ValidationPlugin  string          `json:"validation_plugin"`
	Approvals         map[string]bool `json:"approvals"`
}

type queryRequest struct {
	ChannelName string `json:"channel_name"`
	CCName      string `json:"cc_name"`
}

// QueryApprovedCC
// @Summary Query an approved chaincode definition on a channel.
// @Description `peer lifecycle chaincode queryapproved` is executed through `exec.Command()` to query approved chaincode definitions.
// @Accept json
// @Param body body queryRequest true "cc name and the channel it was approved in"
// @Produce json
// @Tags lifecycle
// @Success 200 {object} approvedChaincodeResponse "successful operation"
// @Router /fabric/lifecycle/approve [get]
func QueryApprovedCC(c *gin.Context) {
	var requestBody queryRequest
	var responseBody approvedChaincodeResponse
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}

	cmd := exec.Command("peer", "lifecycle", "chaincode", "queryapproved", "-C", requestBody.ChannelName,
		"--name", requestBody.CCName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": errMessage})
		return
	}

	outputList := strings.Split(string(output), ":")
	fmt.Println(outputList[2][1:strings.Index(outputList[2], ",")])
	i, _ := strconv.Atoi(outputList[2][1:strings.Index(outputList[2], ",")])
	responseBody.Sequence = int32(i)
	responseBody.Version = outputList[3][1:strings.Index(outputList[3], ",")]
	b, _ := strconv.ParseBool(outputList[4][1:strings.Index(outputList[4], ",")])
	responseBody.InitRequired = b

	c.IndentedJSON(http.StatusOK, string(output))
}

// @Summary Query the committed chaincode definitions by channel on a peer.
// @Description `peer lifecycle chaincode querycommited` is executed through `exec.Command()` to query committed chaincode definitions.
// @Accept json
// @Param body body queryRequest true "cc name and the channel it was committed in"
// @Produce json
// @Tags lifecycle
// @Success 200 {object} committedChaincodeResponse "successful operation"
// @Router /fabric/lifecycle/commit [get]
func queryCommittedCC(c *gin.Context) {

}

// QueryInstalledCC
// @Summary Query the installed chaincodes on a peer.
// @Description `peer lifecycle chaincode queryinstalled` is executed through `exec.Command()` to query installed chaincodes on a peer.
// @Accept json
// @Produce json
// @Tags lifecycle
// @Success 200 {object} installedChaincodeResponse "successful operation"
// @Router /fabric/lifecycle/install [get]
func QueryInstalledCC(c *gin.Context) {
	var response installedChaincodeResponse

	cmd := exec.Command("peer", "lifecycle", "chaincode", "queryinstalled")
	output, err := cmd.CombinedOutput()
	if err != nil {
		errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": errMessage})
		return
	}

	// Installed chaincodes on peer:\nPackage ID: basic_1.0:78f5a4ffe41b97a9615f0c84af8c1dfaa85ce80552494765317ba79c6c15bea1, Label: basic_1.0\n
	outputList := strings.Split(string(output), ":")

	if len(outputList) == 2 { // i.e. Installed chaincodes on peer:
		c.IndentedJSON(http.StatusOK, gin.H{"message": "No chaincode currently installed."})
		return
	}

	// 78f5a4ffe41b97a9615f0c84af8c1dfaa85ce80552494765317ba79c6c15bea1
	packageID := strings.Split(outputList[3], ",")[0]
	// basic_1.0
	label := outputList[4][1 : len(outputList[4])-1]
	response.PackageID = fmt.Sprintf("%s:%s", label, packageID)
	response.Label = label

	c.IndentedJSON(http.StatusOK, response)
}
