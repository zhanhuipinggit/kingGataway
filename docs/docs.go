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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/admin_info": {
            "get": {
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "获取用户信息"
                ],
                "summary": "获取用户信息",
                "operationId": "/admin/admin_info",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminInfoOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/change_pwd": {
            "post": {
                "description": "修改密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "修改密码",
                "operationId": "/admin/change_pwd",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ChangePwdInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin_login/login": {
            "post": {
                "description": "管理员登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员接口"
                ],
                "summary": "管理员登录",
                "operationId": "/admin_login/login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AdminLoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.AdminLoginOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin_login/login_out": {
            "get": {
                "description": "管理员退出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "管理员退出"
                ],
                "summary": "管理员退出",
                "operationId": "/admin_login//login_out",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/demo/bind": {
            "post": {
                "description": "测试数据绑定",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "测试数据绑定",
                "operationId": "/demo/bind",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DemoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.DemoInput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AdminInfoOutput": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "login_time": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "dto.AdminLoginInput": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "user_name": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "dto.AdminLoginOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "dto.ChangePwdInput": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "dto.DemoInput": {
            "type": "object",
            "required": [
                "age",
                "name",
                "passwd"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 20
                },
                "name": {
                    "type": "string",
                    "example": "姓名"
                },
                "passwd": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "middleware.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "errmsg": {
                    "type": "string"
                },
                "errno": {
                    "type": "integer"
                },
                "stack": {
                    "type": "object"
                },
                "trace_id": {
                    "type": "object"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
