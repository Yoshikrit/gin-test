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
            "name": "Walter White",
            "url": "https://twitter.com/example",
            "email": "example@gmail.com"
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
        "/producttype/": {
            "get": {
                "description": "Get all producttype",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Get All ProductType",
                "responses": {
                    "200": {
                        "description": "Get ProductTypes Successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ProductType"
                            }
                        }
                    },
                    "404": {
                        "description": "Error Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    },
                    "500": {
                        "description": "Error Unexpected Error",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create producttype",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Create ProductType",
                "parameters": [
                    {
                        "description": "ProductType data to be create",
                        "name": "ProductType",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProductTypeCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Create ProductType Successfully",
                        "schema": {
                            "$ref": "#/definitions/models.ProductType"
                        }
                    },
                    "400": {
                        "description": "Error Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    },
                    "409": {
                        "description": "Error Conflict Error",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    },
                    "500": {
                        "description": "Error Unexpected Error",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    }
                }
            }
        },
        "/producttype/count": {
            "get": {
                "description": "Get producttype's count from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Get ProductType Count",
                "responses": {
                    "200": {
                        "description": "Get ProductType's Count Successfully",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Error Unexpected Error",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    }
                }
            }
        },
        "/producttype/health": {
            "get": {
                "description": "Health check",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Health Check",
                "operationId": "HealthCheckHandler",
                "responses": {
                    "200": {
                        "description": "ProductType service is running",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/producttype/{id}": {
            "get": {
                "description": "Get producttype by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Get ProductType",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductType ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Get ProductType Successfully",
                        "schema": {
                            "$ref": "#/definitions/models.ProductType"
                        }
                    },
                    "400": {
                        "description": "Error Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    },
                    "404": {
                        "description": "Error Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    },
                    "500": {
                        "description": "Error Unexpected Error",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update producttype by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Update ProductType",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductType ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "ProductType data to be update",
                        "name": "ProductType",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ProductTypeUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Update ProductType Successfully",
                        "schema": {
                            "$ref": "#/definitions/models.ProductType"
                        }
                    },
                    "400": {
                        "description": "Error Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    },
                    "404": {
                        "description": "Error Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    },
                    "500": {
                        "description": "Error Unexpected Error",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete producttype by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "producttype"
                ],
                "summary": "Delete ProductType",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductType ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete ProductType Successfully",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Error Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    },
                    "404": {
                        "description": "Error Not Found",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    },
                    "500": {
                        "description": "Error Unexpected Error",
                        "schema": {
                            "$ref": "#/definitions/errs.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errs.AppError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.ProductType": {
            "type": "object",
            "properties": {
                "ProdType_Id": {
                    "type": "integer"
                },
                "ProdType_Name": {
                    "type": "string"
                }
            }
        },
        "models.ProductTypeCreate": {
            "type": "object",
            "required": [
                "ProdType_Id",
                "ProdType_Name"
            ],
            "properties": {
                "ProdType_Id": {
                    "type": "integer",
                    "minimum": 0
                },
                "ProdType_Name": {
                    "type": "string",
                    "maxLength": 40
                }
            }
        },
        "models.ProductTypeUpdate": {
            "type": "object",
            "required": [
                "ProdType_Name"
            ],
            "properties": {
                "ProdType_Name": {
                    "type": "string",
                    "maxLength": 40
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Gin-Test - ProductType API",
	Description:      "Gin-Test - Teletubbie's ProductType API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
