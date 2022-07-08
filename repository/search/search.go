package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type elasticClient struct {
	es        *elasticsearch.Client
	IndexName string
}

type ESResponse struct {
	Took int64
	Hits struct {
		Total struct {
			Value int64
		}
		Hits []*ESHit
	}
}

type ESHit struct {
	Score   float64 `json:"_score"`
	Index   string  `json:"_index"`
	Type    string  `json:"_type"`
	Version int64   `json:"_version,omitempty"`

	Source Article `json:"_source"`
}

type Article struct {
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	Author          string              `json:"author"`
	UploadDate      string              `json:"uploaded"`
	Description     string              `json:"description"`
	Platform        string              `json:"platform"`
	SignaturePolicy string              `json:"signature_policy"`
	CCLanguages     []*CCLanguage       `json:"cc_languages"`
	AppLanguages    []map[string]string `json:"app_languages"`
	Versions        []map[string]string `json:"versions"`
}

type CCLanguage struct {
	Language     string            `json:"language"`
	Link         string            `json:"link"`
	AssetStruct  map[string]string `json:"asset_struct"`
	Dependencies map[string]string `json:"dependencies"`
}

var (
	esClient = &elasticClient{}
)

func init() {
	cfg := esClientConfig()
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	} else {
		log.Println(elasticsearch.Version)
		log.Println(es.Info())
	}
	esClient.es = es
	esClient.IndexName = "smart_contract"
}

func esClientConfig() elasticsearch.Config {
	cfg := elasticsearch.Config{
		Addresses:              []string{"https://localhost:9200"},
		APIKey:                 "YTQxNTA0RUJnaTRWekEyY2pqOGs6bVF2UFZPSUVSSm1IU0FJcTlxdmNDZw==",
		CertificateFingerprint: "02dfe14e1ae96a59695b5821893380e6dae3e264ca0b27bd176c7b6866f6a5c7",
	}
	return cfg
}

func ESSearchAll(c *gin.Context) {

	searchString := c.Query("filter")
	log.Println(searchString)

	var searchRequest map[string]interface{}
	err := json.Unmarshal([]byte(searchString), &searchRequest)
	if err != nil {
		log.Println(err)
	}

	if searchRequest["q"] == nil {
		searchRequest["q"] = "*"
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": searchRequest["q"],
			},
		},
	}

	res, err := esClient.es.Search(
		esClient.es.Search.WithIndex(esClient.IndexName),
		esClient.es.Search.WithBody(esutil.NewJSONReader(query)),
		esClient.es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer res.Body.Close()

	var sr ESResponse
	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		log.Printf("Error: %s\n", err)
	}

	var scs []Article

	for _, h := range sr.Hits.Hits {
		scs = append(scs, h.Source)
	}
	c.Header("content-range", fmt.Sprintf("%d", len(scs)))

	if scs == nil {
		scs = make([]Article, 0)
	}
	c.IndentedJSON(http.StatusOK, scs)
}

func ESSearchWithLanguage(c *gin.Context) {
	lang := c.Query("language")

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"cc_languages.language": lang,
			},
		},
	}

	res, err := esClient.es.Search(
		esClient.es.Search.WithIndex("smart_contract"),
		esClient.es.Search.WithBody(esutil.NewJSONReader(query)),
		esClient.es.Search.WithPretty(),
	)

	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer res.Body.Close()

	var sr ESResponse
	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		log.Printf("Error: %s\n", err)
	}

	var scs []Article

	for _, h := range sr.Hits.Hits {
		scs = append(scs, h.Source)
	}
	log.Println(scs)
	c.IndentedJSON(http.StatusOK, scs)
}

func EsDocumentByID(c *gin.Context) {
	ccId := c.Param("id")
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": ccId,
			},
		},
	}

	res, err := esClient.es.Search(
		esClient.es.Search.WithIndex("smart_contract"),
		esClient.es.Search.WithBody(esutil.NewJSONReader(query)),
		esClient.es.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer res.Body.Close()

	var sr ESResponse
	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		log.Printf("Error: %s\n", err)
	}

	c.IndentedJSON(http.StatusOK, sr.Hits.Hits[0].Source)
}

func AddDocumentToES(item *Article) (string, error) {
	payload, err := json.Marshal(item)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	res, err := esapi.CreateRequest{
		Index:      esClient.IndexName,
		DocumentID: item.ID,
		Body:       bytes.NewReader(payload),
	}.Do(ctx, esClient.es)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return "", err
		}
		return "", fmt.Errorf("[%s] %s: %s", res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"])
	}

	return "Contract successfully added to search index", nil
}
