package lifecycle

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
)

type installCCRequest struct {
	PackageName string `json:"package_name"`
}

type packageCCRequest struct {
	PackageName string `json:"package_name"`
	Label       string `json:"label"`
	Language    string `json:"language"`
}

// PackageCC
// @Summary Package a cc.
// @Description `peer lifecycle chaincode install` is executed through `exec.Command()` to install chaincode on a peer.
// @Accept json
// @Param body body packageCCRequest true "name of the cc to package (e.g. basic.tar.gz), the language it is written in, and the label for the cc once packaging is done"
// @Produce json
// @Tags lifecycle
// @Success 200 "successful operation"
// @Router /fabric/lifecycle/package [post]
func PackageCC(c *gin.Context) {
	var requestBody packageCCRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}

	c.IndentedJSON(http.StatusOK, requestBody)
}

// InstallCC
// @Summary Install a cc.
// @Description `peer lifecycle chaincode install` is executed through `exec.Command()` to install chaincode on a peer.
// @Accept json
// @Param package_name path string true "name of the package to install (e.g. basic.tar.gz)"
// @Produce json
// @Tags lifecycle
// @Success 200 "successful operation"
// @Router /fabric/lifecycle/install/{package_name} [post]
func InstallCC(c *gin.Context) {

	fileNameParameter := c.Param("package_name")

	fileExists, pathToPackage, err := fileExists(fileNameParameter)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if !fileExists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Package %s does not exist", fileNameParameter)})
		return
	}

	cmd := exec.Command("peer", "lifecycle", "chaincode", "install", pathToPackage)
	output, err := cmd.CombinedOutput()
	if err != nil {
		errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": errMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Package successfully installed."})
}

// fileExists checks if the requested file exists in test-network's directory.
// returns true and the full path of the file
func fileExists(fileName string) (bool, string, error) {
	gopath := os.Getenv("GOPATH")
	path := gopath + "/src/github.com/hyperledger/fabric-samples/test-network/" + fileName
	_, err := os.Stat(path)
	if err == nil {
		return true, path, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, "", nil
	}
	return false, "", err
}
