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
        "/count": {
            "get": {
                "description": "Get number of transactions per day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Count"
                ],
                "summary": "Count daily transactions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/transaction.DailyCount"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpjson.Response"
                        }
                    }
                }
            }
        },
        "/count/all": {
            "get": {
                "description": "Get number of transactions per day, and number of transactions per day per type",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Count"
                ],
                "summary": "Count daily transactions, with count per type",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/server.DailyAllResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpjson.Response"
                        }
                    }
                }
            }
        },
        "/count/type": {
            "get": {
                "description": "Get number of transactions per day per transaction type",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Count"
                ],
                "summary": "Count daily transactions per type",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/transaction.DailyTypeCount"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpjson.Response"
                        }
                    }
                }
            }
        },
        "/max/withdraw": {
            "get": {
                "description": "Get ATM with max withdraw amount per day",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Max"
                ],
                "summary": "ATM with max withdraw per day",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/transaction.DailyMaxWithdraw"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpjson.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpjson.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "server.DailyAllResponse": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/transaction.DailyTypeCount"
                    }
                },
                "total": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/transaction.DailyCount"
                    }
                }
            }
        },
        "transaction.DailyCount": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                }
            }
        },
        "transaction.DailyMaxWithdraw": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "atm_id": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "string"
                }
            }
        },
        "transaction.DailyTypeCount": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "type": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1/daily",
	Schemes:          []string{},
	Title:            "ATM Report API",
	Description:      "ATM report service API server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
