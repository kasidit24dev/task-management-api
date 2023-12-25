# Task management api

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Running the Project](#running-the-project)
* [Technologies](#technologies)
* [API Documentation](#api-documentation)
* [Available API Endpoints](#available-api-endpoints)

## Overview

This API provides a RESTful interface for managing tasks, built in Golang. It adheres to best practices, including
proper error handling, input validation, concurrent access considerations, and comprehensive testing.

### Technologies

**Project is created with:**

* Echo version: 4.11.4
* Viper version: 1.18.2
* Uber-go/zap version: 1.21
* Swaggo/swag version : 1.16.2

### Prerequisites

- Golang (version 1.18 or later recommended)

  [![Golang][Go.dev]][Go-url]

### Installation

1. Clone the repository:

```bash
   git clone https://github.com/kasidit24dev/task-management-api.git
```

2. Install / Update dependencies

```bash
    make update_package
```

### Running the Project

1. Start the server

```bash
    make run
``` 

2. The API will be accessible at ```http://localhost:3000```

### API Documentation

Refer to the `docs/swagger.json` file for Swagger documentation. Access the interactive Swagger UI
at `http://localhost:3000/api/v1/swagger/index.html`.

### Available API Endpoints
BaseUrl: `http://localhost:3000`

| NAME               | METHOD | ENDPOINT         |   REQUIRED   |
|--------------------|:------:|------------------|:------------:|
| Create Task        |  POST  | api/v1/task      |     body     |
| Get task by id     |  GET   | api/v1/task      |    param     |
| Update task        |  PUT   | api/v1/task      | param / body |
| Update task status | PATCH  | api/v1/task      | param / body |
| Delete task        | DELETE | api/v1/task      |    param     |
| Get task list      |  GET   | api/v1/task-list |     N/A      |

[Go.dev]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white

[Go-url]: https://go.dev
