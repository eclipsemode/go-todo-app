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
        "/todos": {
            "get": {
                "description": "get all todos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Show a to-do",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/todos.getAllTodosRes"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "create to-do",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Create to-do post method",
                "parameters": [
                    {
                        "description": "Create to-do req",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/todos.createTodoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/todos.createTodoRes"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "get to-do by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Get To-do by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todos id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/todos.getTodoRes"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "update to-do by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Update To-do by id",
                "parameters": [
                    {
                        "description": "Update to-do req",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/todos.updateTodoReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Todos id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/todos.updateTodoRes"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete to-do by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Delete To-do by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todos id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    },
                    "500": {
                        "description": "Error",
                        "schema": {
                            "$ref": "#/definitions/responseApi.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Todo": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "responseApi.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "todos.createTodoReq": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 100
                },
                "title": {
                    "type": "string",
                    "maxLength": 20
                }
            }
        },
        "todos.createTodoRes": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "todos.getAllTodosRes": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "todos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Todo"
                    }
                }
            }
        },
        "todos.getTodoRes": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "todo": {
                    "$ref": "#/definitions/models.Todo"
                }
            }
        },
        "todos.updateTodoReq": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 100
                },
                "title": {
                    "type": "string",
                    "maxLength": 20
                }
            }
        },
        "todos.updateTodoRes": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "List to do API",
	Description:      "Simple Rest Api for to do list",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
