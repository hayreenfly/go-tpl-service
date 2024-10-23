// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/example": {
            "get": {
                "description": "Returns an example message",
                "tags": [
                    "Example"
                ],
                "summary": "Example page",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JsonResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Returns the health status of the service",
                "tags": [
                    "Health"
                ],
                "summary": "Health check endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.JsonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.JsonResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Use ` + "`" + `interface{}` + "`" + ` for any type of data"
                },
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go Template Service API",
	Description:      "This is a Golang Template Service example with Gin and Swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}