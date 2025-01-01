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
            "email": "erdemkosk@gmail.com"
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
        "/foo": {
            "get": {
                "description": "Get all foo items from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foo"
                ],
                "summary": "Get all items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/model.Foo"
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new foo item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foo"
                ],
                "summary": "Create a new foo",
                "parameters": [
                    {
                        "description": "Foo object",
                        "name": "foo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateFooRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Foo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        },
        "/foo/{id}": {
            "get": {
                "description": "Get a foo item by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foo"
                ],
                "summary": "Get a foo by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Foo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Foo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "Update an existing foo item by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foo"
                ],
                "summary": "Update a foo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Foo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Foo object",
                        "name": "foo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateFooRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{'success': true}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Foo not found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {}
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Get server health status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health check endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateFooRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Foo": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "created_at": {
                    "description": "The creation timestamp",
                    "type": "string"
                },
                "description": {
                    "description": "The description of the foo item",
                    "type": "string",
                    "maxLength": 500,
                    "minLength": 10,
                    "example": "This is an example foo item"
                },
                "id": {
                    "description": "The unique identifier for the foo item",
                    "type": "string",
                    "example": "5f7b5e1b9b0b3a1b3c9b0b3a"
                },
                "name": {
                    "description": "The name of the foo item",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3,
                    "example": "Example Foo"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8090",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Fiber FX Boilerplate API",
	Description:      "This is a sample server for Fiber Boilerplate.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
