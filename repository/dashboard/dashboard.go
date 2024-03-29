package dashboard

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/arogyaGurkha/GurkhaFabricAPI/repository/search"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type installCC struct {
	CCName     string `json:"cc_name"`
	CCPath     string `json:"cc_path"`
	CCLanguage string `json:"cc_language"`
}

type transactionRequest struct {
	AssetID  string `json:"asset_id"`
	NewOwner string `json:"new_owner"`
}

type transactionResponse struct {
	Peer string `json:"peer"`
}

type assetQueryResponse struct {
	Assets []*asset `json:""`
}

type asset struct {
	Value int    `json:"value"`
	Color int    `json:"color"`
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Model string `json:"model"`
	EV    bool   `json:"ev"`
}

// InstallWithDeployCC
// @Summary Install specified CC using deployCC script.
// @Produce json
// @Tags dashboard
// @Success 200 "successful operation"
// @Router /fabric/dashboard/deployCC [post]
func InstallWithDeployCC(c *gin.Context) {

	GOPATH := os.Getenv("GOPATH")
	networkPath := fmt.Sprintf("%s/src/github.com/hyperledger/fabric-samples/test-network", GOPATH)
	ccPathRoot := "../GurkhaContracts"

	var requestBody installCC
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}

	finalCCPath := fmt.Sprintf("%s/%s", ccPathRoot, requestBody.CCPath)

	cmd := exec.Command("bash", "network.sh", "deployCC", "-ccn", requestBody.CCName, "-ccp", finalCCPath, "-ccl", requestBody.CCLanguage)
	cmd.Dir = networkPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": errMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "CC Installed"})
}

// AddDataToES
// @Summary Add document to search index.
// @Description Receive data from UI to upload to the search index. Auto inserts random ID and upload date values.
// @Accept json
// @Param body body search.Article true "Document that needs to be uploaded to the search index."
// @Produce json
// @Tags dashboard
// @Success 200 {object} search.Article
// @Router /fabric/dashboard/smart-contracts [post]
func AddDataToES(c *gin.Context) {
	var searchArticle search.Article
	if err := c.BindJSON(&searchArticle); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	searchArticle.ID = createRandomSHA()
	searchArticle.UploadDate = fmt.Sprintf(time.Now().UTC().Format("2006-01-02"))

	res, err := search.AddDocumentToES(&searchArticle)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func QueryAssets(c *gin.Context) {
	GOPATH := os.Getenv("GOPATH")
	networkPath := fmt.Sprintf("%s/src/github.com/hyperledger/fabric-samples/test-network/scripts", GOPATH)

	cmd := exec.Command("bash", "queryAsset.sh")
	cmd.Dir = networkPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": errMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, string(output))
}

func CreateTransaction(c *gin.Context) {
	var transactionRequest transactionRequest
	if err := c.BindJSON(&transactionRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	//
	//cmd := exec.Command("peer", "version")
	//output, err := cmd.CombinedOutput()
	//if err != nil {
	//	errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
	//	c.IndentedJSON(http.StatusForbidden, gin.H{"message": errMessage})
	//	return
	//}

	c.IndentedJSON(http.StatusOK, transactionRequest)
}

func createRandomSHA() string {
	data := make([]byte, 10)
	var sha string
	if _, err := rand.Read(data); err == nil {
		sha = fmt.Sprintf("%x", sha256.Sum256(data))
	}
	return sha
}
