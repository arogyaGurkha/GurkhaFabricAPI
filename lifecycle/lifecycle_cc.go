package lifecycle

import "github.com/gin-gonic/gin"

type installCCRequest struct {
	PackageName string `json:"package_name"`
}

type packageCCRequest struct {
	PackageName string `json:"package_name"`
	Label       string `json:"label"`
	Language    string `json:"language"`
}

// @Summary Package a cc.
// @Description `peer lifecycle chaincode install` is executed through `exec.Command()` to install chaincode on a peer.
// @Accept json
// @Param body body packageCCRequest true "name of the cc to package (e.g. basic.tar.gz), the language it is written in, and the label for the cc once packaging is done"
// @Produce json
// @Tags lifecycle
// @Success 200 "successful operation"
// @Router /fabric/lifecycle/package [post]
func packageCC(c *gin.Context) {

}

// @Summary Install a cc.
// @Description `peer lifecycle chaincode install` is executed through `exec.Command()` to install chaincode on a peer.
// @Accept json
// @Param package_name path string true "name of the package to install (e.g. basic.tar.gz)"
// @Produce json
// @Tags lifecycle
// @Success 200 "successful operation"
// @Router /fabric/lifecycle/install/{package_name} [post]
func installCC(c *gin.Context) {

}
