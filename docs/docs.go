// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "hertz-contrib",
            "url": "https://github.com/hertz-contrib"
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
        "/blog_server/user/": {
            "get": {
                "responses": {}
            }
        },
        "/blog_server/user/login/": {
            "post": {
                "responses": {}
            }
        },
        "/blog_server/user/profile_modify/": {
            "post": {
                "responses": {}
            }
        },
        "/blog_server/user/pwd_modify/": {
            "post": {
                "responses": {}
            }
        },
        "/blog_server/user/register/": {
            "post": {
                "description": "用户通过提供用户名和密码注册账户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关接口"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户注册请求参数",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserActionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回用户ID及状态信息",
                        "schema": {
                            "$ref": "#/definitions/user.UserActionResponse"
                        }
                    },
                    "400": {
                        "description": "请求参数错误或其他错误信息",
                        "schema": {
                            "$ref": "#/definitions/user.UserActionResponse"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "测试 Description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "测试 Summary",
                "responses": {}
            }
        }
    },
    "definitions": {
        "user.UserActionRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "description": "registered user name",
                    "type": "string"
                }
            }
        },
        "user.UserActionResponse": {
            "type": "object",
            "properties": {
                "status_code": {
                    "type": "integer"
                },
                "status_msg": {
                    "type": "string"
                },
                "user_id": {
                    "description": "user id",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:18005",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "blog_server tests",
	Description:      "This is a demo using Hertz.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}