# Simple Go HTTP Server

A simple HTTP server written in Go that serves JSON responses from two different endpoints.

## Features

- Built using Go's standard `net/http` package
- Returns JSON responses
- Two API endpoints:
  - `/user`
  - `/person`
- Uses appropriate HTTP status codes
- Logs server startup and encoding errors

## Project Structure

```
.
|__ go mod
├── main.go
└── README.md
```

## Requirements

- Go 1.18 or later

## Installation

Clone the repository:

```bash
git clone <repository-url>
```

Navigate into the project directory:

```bash
cd <project-folder>
```

Run the server:

```bash
go run .
```

The server starts on:

```
http://localhost:8080
```

---

## API Endpoints

### GET /user

Returns a user's profile.

**Request**

```http
GET /user HTTP/1.1
Host: localhost:8080
```

**Response**

Status Code:

```
202 Accepted
```

Response Body:

```json
{
    "id": 1,
    "name": "chekus",
    "age": 26,
    "email": "chekusjoseph@gmail.com"
}
```

---

### GET /person

Returns a person's profile.

**Request**

```http
GET /person HTTP/1.1
Host: localhost:8080
```

**Response**

Status Code:

```
202 Accepted
```

Response Body:

```json
{
    "id": 2,
    "name": "joseph",
    "email": "okaforchekus@gmail.com"
}
```

---

## Testing with curl

Retrieve the user profile:

```bash
curl http://localhost:8080/user
```

Retrieve the person profile:

```bash
curl http://localhost:8080/person
```

---

## Technologies Used

- Go
- net/http
- encoding/json

---

## Learning Objectives

This project demonstrates how to:

- Create an HTTP server in Go
- Register routes using `http.NewServeMux`
- Handle HTTP requests
- Return JSON responses
- Set HTTP headers
- Send HTTP status codes
- Encode Go structs into JSON

---

## Future Improvements

- Add POST endpoints
- Validate HTTP methods
- Return `200 OK` for successful GET requests instead of `202 Accepted`
- Add custom error responses
- Refactor repeated JSON response logic into a helper function
- Connect to a database
- Add unit tests
- Implement CRUD operations

---

## Author

**Chekus-dev**
