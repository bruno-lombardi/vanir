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
        "/v1/users": {
            "post": {
                "description": "Creates a new user with the provided information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "create user params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}": {
            "get": {
                "description": "Gets an existent user by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get a user by its ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateUserParams": {
            "description": "Create user params information with email, name and password with confirmation",
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "password_confirmation"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "example": "bruno.lombardi@email.com"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2,
                    "example": "Bruno Lombardi"
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6,
                    "example": "123456"
                },
                "password_confirmation": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6,
                    "example": "123456"
                }
            }
        },
        "models.User": {
            "description": "User account information with user id and email",
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "integer",
                    "example": 1733583441703
                },
                "email": {
                    "type": "string",
                    "example": "bruno.lombardi@email.com"
                },
                "id": {
                    "type": "string",
                    "example": "u_AksOKxc12a"
                },
                "name": {
                    "type": "string",
                    "example": "Bruno Lombardi"
                },
                "updated_at": {
                    "type": "integer",
                    "example": 1733583441710
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
