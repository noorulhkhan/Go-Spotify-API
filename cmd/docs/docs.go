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
        "/track/find/{artist}": {
            "get": {
                "description": "Get tracks collection by Artist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get tracks collection by Artist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Artist",
                        "name": "artist",
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
                                "$ref": "#/definitions/main.Track"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrResp"
                        }
                    }
                }
            }
        },
        "/track/search/{title}": {
            "get": {
                "description": "Gets a track by Title or ISRC",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Gets a track by Title or ISRC",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Track"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ErrResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "main.Artist": {
            "type": "object",
            "properties": {
                "artist_id": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "uri": {
                    "type": "string"
                }
            }
        },
        "main.Image": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "height": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "main.Track": {
            "type": "object",
            "properties": {
                "artists": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Artist"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Image"
                    }
                },
                "isrc": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "track_id": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "utils.ErrResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "record not found"
                },
                "message": {
                    "type": "string",
                    "example": "Record not found"
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
