// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "arogya.Gurkha",
            "url": "https://github.com/arogyaGurkha"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/fabric/lifecycle/admin": {
            "get": {
                "description": "Use terminal environmental variables to get the admin for peer cli container. Only Org1 and Org2 are supported.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lifecycle"
                ],
                "summary": "Get the current admin org.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/lifecycle.currentAdmin"
                        }
                    }
                }
            }
        },
        "/fabric/lifecycle/admin/{organization}": {
            "post": {
                "description": "Use terminal environmental variables to set the admin for peer cli container. Only Org1 and Org2 are supported.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lifecycle"
                ],
                "summary": "Set an org as the admin.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "organization to be set as admin (Org1 and Org2 supported)",
                        "name": "organization",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/lifecycle.currentAdmin"
                        }
                    }
                }
            }
        },
        "/fabric/lifecycle/approve": {
            "get": {
                "description": "` + "`" + `peer lifecycle chaincode queryapproved` + "`" + ` is executed through ` + "`" + `exec.Command()` + "`" + ` to query approved chaincode definitions.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lifecycle"
                ],
                "summary": "Query an approved chaincode definition on a channel.",
                "parameters": [
                    {
                        "description": "cc name and the channel it was approved in",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/lifecycle.queryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/lifecycle.approvedChaincodeResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "` + "`" + `peer lifecycle chaincode approveformyorg` + "`" + ` is executed through ` + "`" + `exec.Command()` + "`" + ` to approve a chaincode definition.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lifecycle"
                ],
                "summary": "Approve the cc definition for the current org.",
                "parameters": [
                    {
                        "description": "channel name (mychannel), cc name (basic), cc version (1.0), cc sequence (1), package ID (run [GET] /fabric/lifecycle/install)",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/lifecycle.approveCCRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation"
                    }
                }
            }
        },
        "/fabric/lifecycle/commit": {
            "get": {
                "description": "` + "`" + `peer lifecycle chaincode querycommited` + "`" + ` is executed through ` + "`" + `exec.Command()` + "`" + ` to query committed chaincode definitions.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lifecycle"
                ],
                "summary": "Query the committed chaincode definitions by channel on a peer.",
                "parameters": [
                    {
                        "description": "cc name and the channel it was committed in",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/lifecycle.queryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/lifecycle.committedChaincodeResponse"
                        }
                    }
                }
            }
        },
        "/fabric/lifecycle/install": {
            "get": {
                "description": "` + "`" + `peer lifecycle chaincode queryinstalled` + "`" + ` is executed through ` + "`" + `exec.Command()` + "`" + ` to query installed chaincodes on a peer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lifecycle"
                ],
                "summary": "Query the installed chaincodes on a peer.",
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/lifecycle.installedChaincodeResponse"
                        }
                    }
                }
            }
        },
        "/fabric/lifecycle/install/{package_name}": {
            "post": {
                "description": "` + "`" + `peer lifecycle chaincode install` + "`" + ` is executed through ` + "`" + `exec.Command()` + "`" + ` to install chaincode on a peer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lifecycle"
                ],
                "summary": "Install a cc.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the package to install (e.g. basic.tar.gz)",
                        "name": "package_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation"
                    }
                }
            }
        },
        "/fabric/lifecycle/package": {
            "post": {
                "description": "` + "`" + `peer lifecycle chaincode install` + "`" + ` is executed through ` + "`" + `exec.Command()` + "`" + ` to install chaincode on a peer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "lifecycle"
                ],
                "summary": "Package a cc.",
                "parameters": [
                    {
                        "description": "name of the cc to package (e.g. asset-transfer-basic), the language it is written in, and the label and package name for the cc once packaging is done",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/lifecycle.packageCCRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation"
                    }
                }
            }
        },
        "/fabric/peer/": {
            "get": {
                "description": "` + "`" + `peer version` + "`" + ` is executed through ` + "`" + `exec.Command()` + "`" + ` to return the current peer version.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "peer"
                ],
                "summary": "Get the current peer binary version",
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/peer.peerVersion"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "lifecycle.approveCCRequest": {
            "type": "object",
            "properties": {
                "cc_name": {
                    "type": "string"
                },
                "cc_sequence": {
                    "type": "integer"
                },
                "cc_version": {
                    "type": "string"
                },
                "channel_name": {
                    "type": "string"
                },
                "package_ID": {
                    "type": "string"
                }
            }
        },
        "lifecycle.approvedChaincodeResponse": {
            "type": "object",
            "properties": {
                "endorsement_plugin": {
                    "type": "string"
                },
                "init_required": {
                    "type": "boolean"
                },
                "package_ID": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                },
                "validation_plugin": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "lifecycle.committedChaincodeResponse": {
            "type": "object",
            "properties": {
                "approvals": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "boolean"
                    }
                },
                "endorsement_plugin": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                },
                "validation_plugin": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "lifecycle.currentAdmin": {
            "type": "object",
            "properties": {
                "admin": {
                    "type": "string"
                }
            }
        },
        "lifecycle.installedChaincodeResponse": {
            "type": "object",
            "properties": {
                "label": {
                    "type": "string"
                },
                "package_ID": {
                    "type": "string"
                }
            }
        },
        "lifecycle.packageCCRequest": {
            "type": "object",
            "properties": {
                "cc_source_name": {
                    "type": "string"
                },
                "label": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "package_name": {
                    "type": "string"
                }
            }
        },
        "lifecycle.queryRequest": {
            "type": "object",
            "properties": {
                "cc_name": {
                    "type": "string"
                },
                "channel_name": {
                    "type": "string"
                }
            }
        },
        "peer.peerVersion": {
            "type": "object",
            "properties": {
                "architecture": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Hyperledger Fabric Gurkhaman API",
	Description: "API to run fabric binaries",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
