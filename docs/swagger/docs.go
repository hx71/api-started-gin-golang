// Package swagger GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

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
        "/api/v1/menus": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menus"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit Per Page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort By {ex: created_at asc | desc}",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Seraching by column {ex: id} action {ex: equals | contains | in}",
                        "name": "id.equals",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menus"
                ],
                "parameters": [
                    {
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/engine.Menu"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/menus/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menus"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menus"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/engine.Menu"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Menus"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/roles": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit Per Page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort By {ex: created_at asc | desc}",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Seraching by column {ex: id} action {ex: equals | contains | in}",
                        "name": "id.equals",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "parameters": [
                    {
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/engine.Role"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/roles/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/engine.Role"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user-menus": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserMenu"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit Per Page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort By {ex: created_at asc | desc}",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Seraching by column {ex: id} action {ex: equals | contains | in}",
                        "name": "id.equals",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserMenu"
                ],
                "parameters": [
                    {
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/engine.UserMenu"
                            }
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user-menus/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserMenu"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserMenu"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/engine.UserMenu"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserMenu"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit Per Page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort By {ex: created_at asc | desc}",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Seraching by column {ex: id} action {ex: equals | contains | in}",
                        "name": "id.equals",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/engine.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "filter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/engine.User"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Pass session information to DBaaS Parameter",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/version": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Api Version"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/engine.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/engine.ResponseStatus"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "engine.Menu": {
            "type": "object",
            "required": [
                "icon",
                "index",
                "main_menu",
                "name",
                "sort",
                "url"
            ],
            "properties": {
                "icon": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "main_menu": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parent": {
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                },
                "sub_parent": {
                    "type": "boolean"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "engine.ResponseStatus": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "engine.ResponseSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "engine.Role": {
            "type": "object",
            "required": [
                "code",
                "name"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "engine.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "engine.UserMenu": {
            "type": "object",
            "required": [
                "is_create",
                "is_delete",
                "is_read",
                "is_report",
                "is_update",
                "menu_id",
                "role_id"
            ],
            "properties": {
                "is_create": {
                    "type": "boolean"
                },
                "is_delete": {
                    "type": "boolean"
                },
                "is_read": {
                    "type": "boolean"
                },
                "is_report": {
                    "type": "boolean"
                },
                "is_update": {
                    "type": "boolean"
                },
                "menu_id": {
                    "type": "string"
                },
                "role_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1234",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Swagger for [Backend API Services]",
	Description:      "This is a document for API use in [Backend API Services]",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
