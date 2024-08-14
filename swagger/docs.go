// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "DevHatt",
            "url": "https://github.com/devhatt"
        },
        "license": {
            "name": "MIT license",
            "url": "https://github.com/devhatt/pet-dex-backend?tab=MIT-1-ov-file#readme"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/pets/": {
            "post": {
                "description": "Sends the Pet's registration data via the request body for persistence in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pet"
                ],
                "summary": "Create Pet by petDto",
                "parameters": [
                    {
                        "description": "Pet object information for registration",
                        "name": "petDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PetInsertDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.PetInsertDto": {
            "type": "object",
            "properties": {
                "adoption_date": {
                    "type": "string",
                    "example": "2008-01-02T15:04:05Z"
                },
                "birthdate": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05Z"
                },
                "breed_id": {
                    "type": "string",
                    "example": "0e0b8399-1bf1-4ed5-a2f4-b5789ddf5df0"
                },
                "name": {
                    "type": "string",
                    "example": "Thor"
                },
                "size": {
                    "type": "string",
                    "example": "medium"
                },
                "user_id": {
                    "type": "string",
                    "example": "fa1b8ae8-5351-11ef-8f02-0242ac130003"
                },
                "weight": {
                    "type": "number",
                    "example": 4.1
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000/api",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "PetDex: Documentação API",
	Description:      "Esta página se destina a documentação da API do projeto PetDex Backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
