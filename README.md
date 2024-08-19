# net-http-guide

This project is a comprehensive guide to using the `net/http` package in Go. It demonstrates various HTTP request handling methods, including GET, POST, PUT, DELETE, query parameters, form handling, URL path parameters, JSON handling, and error handling. The project also includes middleware for logging HTTP requests.

## Table of Contents

- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Server](#running-the-server)
- [Endpoints](#endpoints)
  - [Home](#home)
  - [About](#about)
  - [POST Request](#post-request)
  - [GET Request](#get-request)
  - [PUT Request](#put-request)
  - [Search Query](#search-query)
  - [Form Submission](#form-submission)
  - [Path Parameters](#path-parameters)
  - [JSON Handling](#json-handling)
  - [Error Handling](#error-handling)
  - [Items Management](#items-management)
    - [List All Items](#list-all-items)
    - [Create an Item](#create-an-item)
    - [Get an Item by ID](#get-an-item-by-id)
    - [Update an Item by ID](#update-an-item-by-id)
    - [Delete an Item by ID](#delete-an-item-by-id)
- [Testing](#testing)
- [License](#license)

## Project Structure

```plaintext
net-http-guide/
│
├── directory/
│   ├── main.go
│   ├── main_test.go
│
├── go.mod
├── LICENSE
└── README.md
```
## Getting Started
### Prerequisites

- Go 1.16 or higher installed on your machine.
- Basic understanding of the Go programming language.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/ombima56/net-http-guide.git
cd net-http-guide
```

2. Install dependencies:

```bash
go mod tidy
```
## Running the Server

To start the HTTP server, run the following command:

```go
go run directory/main.go
```

The server will start on port 8080.

## Endpoints
### Home

- URL: /
- Method: GET
- Description: Returns a welcome message.

### About

- URL: /about
- Method: GET
- Description: Returns an about page message.

### POST Request

- URL: /post
- Method: POST
- Description: Handles a POST request.

### GET Request

- URL: /get
- Method: GET
- Description: Handles a GET request.

### PUT Request

- URL: /put
- Method: PUT
- Description: Handles a PUT request.

### Search Query

- URL: /search?q=<query>&sort=<sort>
- Method: GET
- Description: Handles search queries with optional sorting.

### Form Submission

- URL: /form
- Method: GET, POST
- Description: Displays a form for user input and handles form submission.

### Path Parameters

- URL: /items/{id}
- Method: GET, PUT, DELETE
- Description: Handles requests with URL path parameters to manage items by ID.

### JSON Handling

- URL: /items
- Method: GET, POST
- Description: Handles JSON requests to create or list items.

### Error Handling

- URL: /error
- Method: GET
- Description: Returns a custom error message.

## Items Management

These endpoints allow you to manage items with the following operations:
### List All Items

- URL: /items
- Method: GET
- Description: Returns a list of all items.

```bash
curl -X GET http://localhost:8080/items
```
### Create an Item

- URL: /items
- Method: POST
- Description: Creates a new item. Requires a JSON body with name and price.

```bash
curl -X POST http://localhost:8080/items -d '{"name":"Item 1", "price":10.99}' -H "Content-Type: application/json"
```
### Get an Item by ID

- URL: /items/{id}
- Method: GET
- Description: Retrieves an item by its ID.

```bash
curl -X GET http://localhost:8080/items/1
```

### Update an Item by ID

- URL: /items/{id}
- Method: PUT
- Description: Updates an item by its ID. Requires a JSON body with name and/or price.

```bash
curl -X PUT http://localhost:8080/items/1 -d '{"price":15.99}' -H "Content-Type: application/json"
```

### Delete an Item by ID

- URL: /items/{id}
- Method: DELETE
- Description: Deletes an item by its ID.

## Testing

You can run the tests included in the main_test.go file using the following command:

```bash
go test -v.
```

This will execute all the test cases and display the results.

## License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.