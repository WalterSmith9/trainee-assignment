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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Тестовое задание для стажёра Backend",
    "title": "trainee-assignment",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "vova.kuznetsov4@mail.ru"
    },
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/api/v2",
  "paths": {
    "/segment": {
      "post": {
        "description": "Accepts information about segment and stores it",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Segment"
        ],
        "summary": "Create new segment",
        "operationId": "SegmentCreate",
        "parameters": [
          {
            "description": "Segment information fields",
            "name": "segmentInfo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SegmentCreateRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "delete": {
        "description": "Accepts segment slug and removes it",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Segment"
        ],
        "summary": "Removes existing segment",
        "operationId": "SegmentDelete",
        "parameters": [
          {
            "description": "Segment information fields",
            "name": "segmentInfo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SegmentDeleteRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "description": "Accepts user's ID and returns segments they participate in",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "summary": "Return user's segments",
        "operationId": "GetSegments",
        "parameters": [
          {
            "description": "User information fields",
            "name": "userInfo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserSegmentsGetRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/UserSegmentsGetResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "post": {
        "description": "Accepts information about user and segments they participate in and stores it",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "summary": "Update user's segments",
        "operationId": "UserEdit",
        "parameters": [
          {
            "description": "User information fields",
            "name": "userInfo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserEditRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "SegmentCreateRequest": {
      "type": "object",
      "required": [
        "slug"
      ],
      "properties": {
        "slug": {
          "description": "Segment slug",
          "type": "string",
          "example": "AVITO_VOICE_MESSAGES"
        }
      }
    },
    "SegmentDeleteRequest": {
      "type": "object",
      "required": [
        "slug"
      ],
      "properties": {
        "slug": {
          "description": "Segment slug",
          "type": "string",
          "example": "AVITO_VOICE_MESSAGES"
        }
      }
    },
    "SuccessResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "UserEditRequest": {
      "type": "object",
      "required": [
        "user_id"
      ],
      "properties": {
        "slugs_add": {
          "description": "Array of segment slugs to add",
          "type": "array",
          "items": {
            "description": "Segment slug",
            "type": "string",
            "example": "AVITO_VOICE_MESSAGES"
          }
        },
        "slugs_delete": {
          "description": "Array of segment slugs to delete",
          "type": "array",
          "items": {
            "description": "Segment slug",
            "type": "string",
            "example": "AVITO_DISCOUNT_50"
          }
        },
        "user_id": {
          "description": "User id",
          "type": "integer",
          "example": 1002
        }
      }
    },
    "UserSegmentsGetRequest": {
      "type": "object",
      "required": [
        "user_id"
      ],
      "properties": {
        "user_id": {
          "description": "User id",
          "type": "integer",
          "example": 1002
        }
      }
    },
    "UserSegmentsGetResponse": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Response code",
          "type": "integer",
          "format": "int32",
          "example": 1
        },
        "data": {
          "description": "Array of Segments User is in",
          "type": "array",
          "items": {
            "description": "Response data object",
            "type": "string",
            "example": "AVITO_VOICE_MESSAGES"
          }
        },
        "message": {
          "description": "Response message",
          "type": "string",
          "example": "success"
        }
      }
    }
  },
  "tags": [
    {
      "description": "API operations with segment",
      "name": "Segment",
      "externalDocs": {
        "description": "Find out more about segment",
        "url": "http://localhost/swagger/#"
      }
    },
    {
      "description": "API operations with user",
      "name": "User",
      "externalDocs": {
        "description": "Find out more about user",
        "url": "http://localhost/swagger/#"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://localhost/swagger/#"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Тестовое задание для стажёра Backend",
    "title": "trainee-assignment",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "vova.kuznetsov4@mail.ru"
    },
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/api/v2",
  "paths": {
    "/segment": {
      "post": {
        "description": "Accepts information about segment and stores it",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Segment"
        ],
        "summary": "Create new segment",
        "operationId": "SegmentCreate",
        "parameters": [
          {
            "description": "Segment information fields",
            "name": "segmentInfo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SegmentCreateRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "delete": {
        "description": "Accepts segment slug and removes it",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Segment"
        ],
        "summary": "Removes existing segment",
        "operationId": "SegmentDelete",
        "parameters": [
          {
            "description": "Segment information fields",
            "name": "segmentInfo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SegmentDeleteRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "description": "Accepts user's ID and returns segments they participate in",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "summary": "Return user's segments",
        "operationId": "GetSegments",
        "parameters": [
          {
            "description": "User information fields",
            "name": "userInfo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserSegmentsGetRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/UserSegmentsGetResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      },
      "post": {
        "description": "Accepts information about user and segments they participate in and stores it",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "User"
        ],
        "summary": "Update user's segments",
        "operationId": "UserEdit",
        "parameters": [
          {
            "description": "User information fields",
            "name": "userInfo",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserEditRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "SegmentCreateRequest": {
      "type": "object",
      "required": [
        "slug"
      ],
      "properties": {
        "slug": {
          "description": "Segment slug",
          "type": "string",
          "example": "AVITO_VOICE_MESSAGES"
        }
      }
    },
    "SegmentDeleteRequest": {
      "type": "object",
      "required": [
        "slug"
      ],
      "properties": {
        "slug": {
          "description": "Segment slug",
          "type": "string",
          "example": "AVITO_VOICE_MESSAGES"
        }
      }
    },
    "SuccessResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "UserEditRequest": {
      "type": "object",
      "required": [
        "user_id"
      ],
      "properties": {
        "slugs_add": {
          "description": "Array of segment slugs to add",
          "type": "array",
          "items": {
            "description": "Segment slug",
            "type": "string",
            "example": "AVITO_VOICE_MESSAGES"
          }
        },
        "slugs_delete": {
          "description": "Array of segment slugs to delete",
          "type": "array",
          "items": {
            "description": "Segment slug",
            "type": "string",
            "example": "AVITO_DISCOUNT_50"
          }
        },
        "user_id": {
          "description": "User id",
          "type": "integer",
          "example": 1002
        }
      }
    },
    "UserSegmentsGetRequest": {
      "type": "object",
      "required": [
        "user_id"
      ],
      "properties": {
        "user_id": {
          "description": "User id",
          "type": "integer",
          "example": 1002
        }
      }
    },
    "UserSegmentsGetResponse": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Response code",
          "type": "integer",
          "format": "int32",
          "example": 1
        },
        "data": {
          "description": "Array of Segments User is in",
          "type": "array",
          "items": {
            "description": "Response data object",
            "type": "string",
            "example": "AVITO_VOICE_MESSAGES"
          }
        },
        "message": {
          "description": "Response message",
          "type": "string",
          "example": "success"
        }
      }
    }
  },
  "tags": [
    {
      "description": "API operations with segment",
      "name": "Segment",
      "externalDocs": {
        "description": "Find out more about segment",
        "url": "http://localhost/swagger/#"
      }
    },
    {
      "description": "API operations with user",
      "name": "User",
      "externalDocs": {
        "description": "Find out more about user",
        "url": "http://localhost/swagger/#"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://localhost/swagger/#"
  }
}`))
}
