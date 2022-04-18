// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Shivani User Example1",
    "version": "0.1.0"
  },
  "paths": {
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "find_Users",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 0,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the User operations",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "users"
        ],
        "operationId": "add_User",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Name should be atleast 3 characters long",
            "schema": {
              "$ref": "#/definitions/error"
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
    "/users/{id}": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "view_User",
        "responses": {
          "200": {
            "description": "list the particular User operations",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "user is not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "users"
        ],
        "operationId": "delete_user",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "404": {
            "description": "user is not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "address": {
          "description": "Address",
          "type": "string",
          "example": "ABC"
        },
        "created_at": {
          "description": "Timestamp when the user was created",
          "type": "string",
          "format": "date-time",
          "title": "Created At",
          "readOnly": true
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Shivani User Example1",
    "version": "0.1.0"
  },
  "paths": {
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "find_Users",
        "parameters": [
          {
            "type": "string",
            "name": "name",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 0,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the User operations",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "users"
        ],
        "operationId": "add_User",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Name should be atleast 3 characters long",
            "schema": {
              "$ref": "#/definitions/error"
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
    "/users/{id}": {
      "get": {
        "tags": [
          "users"
        ],
        "operationId": "view_User",
        "responses": {
          "200": {
            "description": "list the particular User operations",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "user is not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "users"
        ],
        "operationId": "delete_user",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "404": {
            "description": "user is not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "address": {
          "description": "Address",
          "type": "string",
          "example": "ABC"
        },
        "created_at": {
          "description": "Timestamp when the user was created",
          "type": "string",
          "format": "date-time",
          "title": "Created At",
          "readOnly": true
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
}
