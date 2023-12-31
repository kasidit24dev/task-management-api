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
        "/task": {
            "post": {
                "description": "Create task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreateTask",
                "operationId": "post-task",
                "parameters": [
                    {
                        "description": "Request Body for create task",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TaskRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success\" example={\"code\":201,\"message\":\"Success\"}",
                        "schema": {
                            "$ref": "#/definitions/handlers.commonResponse"
                        }
                    }
                }
            }
        },
        "/task/{id}": {
            "get": {
                "description": "Get a task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "GetTaskByID",
                "operationId": "get-task-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success\" example={\"code\":200,\"message\":\"Success\",\"data\":{\"id\":2,\"title\":\"2\",\"description\":\"xxx\",\"status\":\"To Do\"}}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.commonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.TaskResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "Update a task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "UpdateTask",
                "operationId": "update-task-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request Body for updating task",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success\" example={\"code\":200,\"message\":\"Success\"}",
                        "schema": {
                            "$ref": "#/definitions/handlers.commonResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a task by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "DeleteTaskByID",
                "operationId": "delete-task-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success\" example={\"code\":200,\"message\":\"Success\"}",
                        "schema": {
                            "$ref": "#/definitions/handlers.commonResponse"
                        }
                    }
                }
            }
        },
        "/task/{id}/status": {
            "patch": {
                "description": "Update the status of a task by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "UpdateTaskStatus",
                "operationId": "update-task-status-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New task status [To do, In Progress, Done]",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateStatusRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success\" example={\"code\":200,\"message\":\"Success\"}",
                        "schema": {
                            "$ref": "#/definitions/handlers.commonResponse"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Retrieve a list of tasks",
                "produces": [
                    "application/json"
                ],
                "summary": "ListTasks",
                "operationId": "list-tasks",
                "responses": {
                    "200": {
                        "description": "Success\" example={\"code\":200,\"message\":\"Success\",\"data\":[{\"id\":1,\"title\":\"2\",\"description\":\"xxx\",\"status\":\"To Do\"},{\"id\":2,\"title\":\"2\",\"description\":\"xxx\",\"status\":\"To Do\"}]}",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handlers.commonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.TaskResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ErrData": {
            "type": "object",
            "properties": {
                "error_message": {
                    "type": "string"
                }
            }
        },
        "handlers.commonResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "error": {
                    "$ref": "#/definitions/handlers.ErrData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.TaskRequest": {
            "type": "object",
            "required": [
                "description",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.TaskResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.UpdateStatusRequest": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Example Response for Task Management API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
