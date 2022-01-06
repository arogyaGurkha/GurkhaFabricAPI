package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"net/http"
	"os"
)

type cloneRequest struct {
	Url       string `json:"url"`
	Directory string `json:"directory"`
}

func AddRemote(c *gin.Context) {
	// git remote add origin git@3.34.46.252:remote_fabric.git
}

// CloneCC
// @Summary Clone a repository.
// @Description Clone a repository.
// @Accept json
// @Param body body cloneRequest true "url (https://github.com/arogyaGurkha/GurkhaContracts.git), directory (~/)"
// @Produce json
// @Tags repository
// @Success 200 "successful operation"
// @Router /fabric/repository/clone [post]
func CloneCC(c *gin.Context) {
	var requestBody cloneRequest
	GOPATH := os.Getenv("GOPATH")
	//rootPath := fmt.Sprintf("%s/src/github.com/hyperledger/fabric-samples/", GOPATH)

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}

	requestBody.Directory = fmt.Sprintf("%s/%s", GOPATH, requestBody.Directory)

	r, err := git.PlainClone(requestBody.Directory, false, &git.CloneOptions{
		URL:               requestBody.Url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": commit})
}

func AddChanges(c *gin.Context) {
	// git add .
}

func CommitChanges(c *gin.Context) {
	// git commit
}

func PushChanges(c *gin.Context) {
	// git push origin master
}
