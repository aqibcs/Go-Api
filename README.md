# Go-Api
Create an API in Go using Gin frame work
# Simple Book Management API

A basic Go web application using the Gin framework to manage a list of books.

## Introduction

This project is a simple Go web application that serves as a RESTful API for managing a list of books. It utilizes the Gin web framework to handle HTTP requests and provides endpoints for actions such as retrieving all books, getting a book by ID, creating a new book, checking out a book, and returning a book.

## Installation

To run the project locally, follow these steps:

1. Clone the repository:

    ```bash
    git clone git@github.com:aqibcs/Go-Api.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Go-Api
    ```

3. Run the Go application:

    ```bash
    go run main.go
    ```

The server will start running at `http://localhost:8080`.

## Usage

The API provides various endpoints to interact with the book data. You can use tools like `curl` or Postman to send HTTP requests. For example:

- Retrieve all books:

    ```bash
    curl http://localhost:8080/books
    ```

- Retrieve a book by ID:

    ```bash
    curl http://localhost:8080/books/1
    ```

- Create a new book:

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"id": "4", "title": "New Book", "author": "Author Name", "quantity": 1}' http://localhost:8080/books
    ```

... (and so on for other endpoints)

## Endpoints

- **GET /books**: Retrieve all books.
- **GET /books/:id**: Retrieve a book by ID.
- **POST /books**: Create a new book.
- **PATCH /books/checkout?id=:id**: Check out a book.
- **PATCH /books/return?id=:id**: Return a book.
