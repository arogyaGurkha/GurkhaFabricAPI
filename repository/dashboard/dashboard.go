package dashboard

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/arogyaGurkha/GurkhaFabricAPI/admin"
	"github.com/arogyaGurkha/GurkhaFabricAPI/repository/search"
	"github.com/gin-gonic/gin"
	"log"
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
	Assets []*Asset `json:""`
}

type Asset struct {
	ID             string `json:"ID"`
	Color          string `json:"color"`
	Size           int    `json:"size"`
	Owner          string `json:"owner"`
	AppraisedValue int    `json:"appraisedValue"`
}

var (
	peerAddressOrg1 = "localhost:7051"
	peerAddressOrg2 = "localhost:9051"
	GOPATH          = os.Getenv("GOPATH")
	networkPath     = fmt.Sprintf("%s/src/github.com/hyperledger/fabric-samples/test-network", GOPATH)
	now             = time.Now()
	assetId         = fmt.Sprintf("asset%d", now.Unix()*1e3+int64(now.Nanosecond()/1e6))
)

// InstallWithDeployCC
// @Summary Install specified CC using deployCC script.
// @Produce json
// @Tags dashboard
// @Success 200 "successful operation"
// @Router /fabric/dashboard/deployCC [post]
func InstallWithDeployCC(c *gin.Context) {
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
	networkScriptPath := fmt.Sprintf("%s/src/github.com/hyperledger/fabric-samples/test-network/scripts", GOPATH)

	cmd := exec.Command("bash", "queryAsset.sh")
	cmd.Dir = networkScriptPath

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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message1": err})
		return
	}
	log.Println(transactionRequest)

	ordererIP := "localhost:7050"
	ordererName := "orderer.example.com"
	ordererCertPath := fmt.Sprintf("%s/organizations/ordererOrganizations/example.com/orderers/"+
		"orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem", networkPath)

	os.Setenv("CORE_PEER_TLS_ENABLED", "true")
	os.Setenv("CORE_PEER_LOCALMSPID", "Org1MSP")
	os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE",
		fmt.Sprintf("%s/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt", networkPath))
	os.Setenv("CORE_PEER_ADDRESS", "localhost:7051")
	os.Setenv("CORE_PEER_MSPCONFIGPATH", fmt.Sprintf("%s/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp", networkPath))

	cmd := exec.Command("peer", "chaincode", "invoke", "-C", "mychannel", "-n", "basic", "-o", ordererIP, "--ordererTLSHostnameOverride", ordererName,
		"--tls", "true", "--cafile", ordererCertPath, "--peerAddresses", peerAddressOrg1, "--tlsRootCertFiles",
		fmt.Sprintf("%s/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt", networkPath),
		"--peerAddresses", peerAddressOrg2, "--tlsRootCertFiles",
		fmt.Sprintf("%s/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt", networkPath),
		"-c", fmt.Sprintf(`{"function":"CreateAsset","Args":["%s","black","10","%s","11"]}`, transactionRequest.AssetID, transactionRequest.NewOwner))
	cmd.Dir = networkPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
		log.Println("err :" + errMessage)
		c.IndentedJSON(http.StatusForbidden, gin.H{"message2": errMessage})
		return
	}
	log.Println(string(output))
	c.IndentedJSON(http.StatusOK, string(output))
}
func UpdateTransaction(c *gin.Context) {
	var transactionRequest transactionRequest
	if err := c.BindJSON(&transactionRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message1": err})
		return
	}
	log.Println(transactionRequest)

	ordererIP := "localhost:7050"
	ordererName := "orderer.example.com"
	ordererCertPath := fmt.Sprintf("%s/organizations/ordererOrganizations/example.com/orderers/"+
		"orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem", networkPath)

	os.Setenv("CORE_PEER_TLS_ENABLED", "true")
	os.Setenv("CORE_PEER_LOCALMSPID", "Org1MSP")
	os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE",
		fmt.Sprintf(
			"%s/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt",
			networkPath))
	os.Setenv("CORE_PEER_ADDRESS", "localhost:7051")
	os.Setenv("CORE_PEER_MSPCONFIGPATH", fmt.Sprintf("%s/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp", networkPath))

	//peer chaincode query  -C mychannel -n basic -c '{"function":"UpdateAsset","Args":["asset1","black","10","jeho","11"]}'
	cmd := exec.Command("peer", "chaincode", "invoke", "-C", "mychannel", "-n", "basic", "-o", ordererIP, "--ordererTLSHostnameOverride", ordererName,
		"--tls", "true", "--cafile", ordererCertPath, "--peerAddresses", peerAddressOrg1, "--tlsRootCertFiles",
		fmt.Sprintf("%s/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt", networkPath),
		"--peerAddresses", peerAddressOrg2, "--tlsRootCertFiles",
		fmt.Sprintf("%s/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt", networkPath),
		"-c", fmt.Sprintf(`{"function":"UpdateAsset","Args":["%s","black","10","%s","11"]}`, transactionRequest.AssetID, transactionRequest.NewOwner))
	cmd.Dir = networkPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
		log.Println("err :" + errMessage)
		c.IndentedJSON(http.StatusForbidden, gin.H{"message2": errMessage})
		return
	}
	log.Println(string(output))
	c.IndentedJSON(http.StatusOK, string(output))
}
func AssetTransfer2(c *gin.Context) {
	var transactionRequest transactionRequest
	if err := c.BindJSON(&transactionRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message1": err})
		return
	}
	log.Println(admin.ContractPass)
	res := admin.TransferAsset(admin.ContractPass, transactionRequest.AssetID, transactionRequest.NewOwner)
	c.IndentedJSON(http.StatusOK, string(res))
	log.Println(res)
}

func createRandomSHA() string {
	data := make([]byte, 10)
	var sha string
	if _, err := rand.Read(data); err == nil {
		sha = fmt.Sprintf("%x", sha256.Sum256(data))
	}
	return sha
}
