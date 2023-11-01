// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/testapi/update-product/{product_id}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "Update product attributes",
                "operationId": "update-product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ProductUpdates"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "main.ProductUpdates": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "product_info.swagger.io",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server for updating product information.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
