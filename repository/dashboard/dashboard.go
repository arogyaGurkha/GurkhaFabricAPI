package dashboard

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/arogyaGurkha/GurkhaFabricAPI/admin"
	"github.com/arogyaGurkha/GurkhaFabricAPI/repository/search"
	"github.com/gin-gonic/gin"
)

type installCC struct {
	CCName     string `json:"cc_name"`
	CCPath     string `json:"cc_path"`
	CCLanguage string `json:"cc_language"`
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
	scriptPath      = fmt.Sprintf("%s/src/github.com/hyperledger/fabric-samples/test-network/scripts", GOPATH)
	now             = time.Now()
	assetId         = fmt.Sprintf("asset%d", now.Unix()*1e3+int64(now.Nanosecond()/1e6))
)

// FileUpload
// @Summary If clients send cc package file, then upload zip file at /Downloads/chaincodes and install system channel.
// @Produce json
// @Tags dashboard
// @Success 200 "successful operation"
// @Router /fabric/dashboard/smart-contracts/file [post]
func FileUpload(c *gin.Context) {
	// get zip file and upload it
	rawData := c.PostForm("data")
	var inputData search.Article
	json.Unmarshal([]byte(rawData), &inputData)
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, fmt.Sprintf(`/home/jeho/Downloads/chaincodes/%s.tar.gz`, inputData.Name))
	log.Println("zip file uploaded successfully")

	// save smart constarct info
	inputData.UploadDate = fmt.Sprintf(time.Now().UTC().Format("2006-01-02"))
	log.Println(inputData.UploadDate)
	res, err := search.AddDocumentToES(&inputData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

// InstallWithDeployCC
// @Summary Install specified CC using deployCC script.
// @Produce json
// @Tags dashboard
// @Success 200 "successful operation"
// @Router /fabric/dashboard/deployCC [post]
func InstallWithDeployCC(c *gin.Context) {
	ccPathRoot := "/home/jeho/Downloads/chaincodes"

	var requestBody installCC
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}

	setEnv()
	finalCCPath := fmt.Sprintf("%s/%s", ccPathRoot, requestBody.CCPath)
	log.Println(fmt.Sprintf("CCName :%s, ccPath :%s, finalCCPath : %s", requestBody.CCName, requestBody.CCPath, finalCCPath))
	cmd := exec.Command("bash", "./scripts/deployTarCC.sh", "mychannel", requestBody.CCName, requestBody.CCLanguage)
	cmd.Dir = networkPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
		log.Println(errMessage)
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
	if err := c.ShouldBindJSON(&searchArticle); err != nil {
		log.Println("add data err")
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	searchArticle.UploadDate = fmt.Sprintf(time.Now().UTC().Format("2006-01-02"))
	log.Println(searchArticle.UploadDate)
	res, err := search.AddDocumentToES(&searchArticle)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func QueryAssets(c *gin.Context) {

	cmd := exec.Command("bash", "queryAsset.sh")
	cmd.Dir = scriptPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(output)
		log.Println(err)
		errMessage := fmt.Sprintf(fmt.Sprint(err) + ": " + string(output))
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": errMessage})
		return
	}
	c.IndentedJSON(http.StatusOK, string(output))
}

func AssetTransfer(c *gin.Context) {
	var transactionRequest admin.TransactionRequest
	if err := c.BindJSON(&transactionRequest); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message1": err})
		return
	}
	log.Println(transactionRequest)

	function := transactionRequest.Function

	switch function {
	case "CreateAsset":
		admin.CreateAsset(admin.ContractPass, transactionRequest)
		log.Println("CreateAsset")
	case "UpdateAsset":
		admin.UpdateAsset(admin.ContractPass, transactionRequest)
		log.Println("UpdateAsset")
	case "TransferAsset":
		admin.TransferAsset(admin.ContractPass, transactionRequest)
		log.Println("TransferAsset")
	default:
		log.Println(function)
		log.Fatalln("function selection error")
	}
}

func createRandomSHA() string {
	data := make([]byte, 10)
	var sha string
	if _, err := rand.Read(data); err == nil {
		sha = fmt.Sprintf("%x", sha256.Sum256(data))
	}
	return sha
}

func setEnv() {
	os.Setenv("CORE_PEER_TLS_ENABLED", "true")
	os.Setenv("CORE_PEER_LOCALMSPID", "Org1MSP")
	os.Setenv("CORE_PEER_TLS_ROOTCERT_FILE",
		fmt.Sprintf("%s/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt", networkPath))
	os.Setenv("CORE_PEER_ADDRESS", "localhost:7051")
	os.Setenv("CORE_PEER_MSPCONFIGPATH", fmt.Sprintf("%s/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp", networkPath))
}
