// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "trisnm102@gmail.com"
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
        "/admin/employees": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create New Employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "Create New Employee",
                "parameters": [
                    {
                        "description": "Create an employee",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Employee"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/v1/employees/{employee_id}/actions": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "User Creates Action",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "User Creates Action",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee ID",
                        "name": "employee_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "employee creates an action",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Action"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Action"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/v1/employees/{employee_id}/attendances": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "User Gets Attendance List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "employee"
                ],
                "summary": "User Gets Attendance List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee ID",
                        "name": "employee_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Attendance"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Action": {
            "type": "object",
            "required": [
                "place_id"
            ],
            "properties": {
                "employee_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "place_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "checkin",
                        "checkout"
                    ]
                }
            }
        },
        "entity.Attendance": {
            "type": "object",
            "required": [
                "employee_id"
            ],
            "properties": {
                "check_in_time": {
                    "type": "string"
                },
                "check_out_time": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "employee_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "over_time": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "total_hour": {
                    "type": "string"
                }
            }
        },
        "entity.Employee": {
            "type": "object",
            "properties": {
                "department": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Interview Be Earning API",
	Description:      "A employee attandance management service API in Go using Gin framework..",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
