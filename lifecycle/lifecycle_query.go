package lifecycle

import "github.com/gin-gonic/gin"

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

// @Summary Query an org's approved chaincode definition from its peer.
// @Description `peer lifecycle chaincode queryapproved` is executed through `exec.Command()` to query approved chaincode definitions.
// @Accept json
// @Param body body queryRequest true "cc name and the channel it was approved in"
// @Produce json
// @Tags lifecycle
// @Success 200 {object} approvedChaincodeResponse "successful operation"
// @Router /fabric/lifecycle/approve [get]
func queryApproved(c *gin.Context) {

}

// @Summary Query the committed chaincode definitions by channel on a peer.
// @Description `peer lifecycle chaincode querycommited` is executed through `exec.Command()` to query committed chaincode definitions.
// @Accept json
// @Param body body queryRequest true "cc name and the channel it was committed in"
// @Produce json
// @Tags lifecycle
// @Success 200 {object} committedChaincodeResponse "successful operation"
// @Router /fabric/lifecycle/commit [get]
func queryCommitted(c *gin.Context) {

}

// @Summary Query the installed chaincodes on a peer.
// @Description `peer lifecycle chaincode queryinstalled` is executed through `exec.Command()` to query installed chaincodes on a peer.
// @Accept json
// @Produce json
// @Tags lifecycle
// @Success 200 {object} installedChaincodeResponse "successful operation"
// @Router /fabric/lifecycle/install [get]
func queryInstalled(c *gin.Context) {

}
