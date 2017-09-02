// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/joinc.user-address.v1+json"
  ],
  "produces": [
    "application/joinc.user-address.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Simple Address API",
    "version": "0.1.0"
  },
  "paths": {
    "/search": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "user_id로 유저를 검색한다.",
        "parameters": [
          {
            "description": "user id",
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/user/{id}": {
      "get": {
        "tags": [
          "user"
        ],
        "summary": "유저 정보를 읽어온다.",
        "parameters": [
          {
            "type": "string",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "integer"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "user_id"
      ],
      "properties": {
        "address": {
          "description": "user address",
          "type": "string"
        },
        "email": {
          "description": "user email",
          "type": "string"
        },
        "name": {
          "description": "user name",
          "type": "string"
        },
        "user_id": {
          "description": "user id",
          "type": "string"
        }
      }
    }
  }
}`))
}
