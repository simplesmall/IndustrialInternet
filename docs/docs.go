// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://www.topgoer.com",
        "contact": {
            "name": "www.topgoer.com",
            "url": "https://www.topgoer.com",
            "email": "example@some.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/login": {
            "get": {
                "description": "普通GET测试接口描述信息",
                "tags": [
                    "Test"
                ],
                "summary": "普通GET测试",
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "$ref": "#/definitions/Response.ResponseStructure"
                        }
                    }
                }
            }
        },
        "/api/msgOK": {
            "get": {
                "description": "NormalOKTest测试接口描述信息",
                "tags": [
                    "Test"
                ],
                "summary": "NormalOKTest测试接口",
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "$ref": "#/definitions/Response.ResponseStructure"
                        }
                    }
                }
            }
        },
        "/api/normal": {
            "get": {
                "description": "NormalTest测试接口描述信息",
                "tags": [
                    "Test"
                ],
                "summary": "NormalTest测试接口",
                "responses": {
                    "200": {
                        "description": "code\":200,\"data\":null,\"msg\":\"\"}",
                        "schema": {
                            "$ref": "#/definitions/Response.ResponseStructure"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "description": "登录测试接口描述信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录"
                ],
                "summary": "登录测试接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"msg\":\"\",\"data\":null}",
                        "schema": {
                            "$ref": "#/definitions/Response.ResponseStructure"
                        }
                    }
                }
            }
        },
        "/api/v1/notfound": {
            "get": {
                "description": "notfound测试接口描述信息",
                "tags": [
                    "NotFound"
                ],
                "summary": "notfound测试接口",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Response.ResponseStructure": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
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
	Version:     "1.0",
	Host:        "localhost:8090",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "IndustrialInternet项目API",
	Description: "This is a  server of IndustrialInternet.",
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
	swag.Register(swag.Name, &s{})
}
