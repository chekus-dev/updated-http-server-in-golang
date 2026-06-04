# Simple Go JSON API Server

Hey there! This is a dynamic, lightweight HTTP server written in Go that validates and serves user data in JSON format instead of traditional HTML. It includes a mock in-memory database lookup and robust automated unit testing.

---

## Features

* **Zero Third-Party Dependencies**: Built entirely using Go's standard library (`net/http`, `encoding/json`, and `strconv`).
* **Dynamic Query Parameter Lookup**: Parses incoming requests to dynamically fetch specific users via URL flags (e.g., `/user?id=1`).
* **Robust Error Handling**: Safely returns semantic HTTP status codes (`400 Bad Request`, `404 Not Found`) along with descriptive JSON error payloads.
* **Streamlined Encoding**: Uses `json.NewEncoder` to stream data directly into the response writer, keeping memory usage ultra-low.
* **Table-Driven Automated Tests**: Fully covered by memory-efficient HTTP unit tests using `net/http/httptest`.

---

## Project Structure

The project maps data models into native Go structs before converting them to raw JSON:

* **`User`**: A custom Go structure specifying the shape of our data schema (`ID`, `Name`, `Email`).
* **`json:"name"`**: Struct tags ensure field keys map to standardized lowercase attributes in the final JSON output.
* **`userDatabase`**: An in-memory Go `map` that simulates database records for instant reading.

---

## How to Run It

### 1. Run the server
Make sure you have Go installed on your machine. Save the code into a file named `main.go` and run the following command in your terminal:

```bash
go run main.go
```

You should see this confirmation logging message in your terminal:
> `starting server on localhost:8080`

### 2. Test the API responses
While the server is running, open a separate terminal window and query the server using `curl` to target different IDs:

#### Fetching User 1:
```bash
curl "http://localhost:8080/user?id=1"
```
**Expected JSON Output:**
```json
{
  "id": 1,
  "name": "chekus joseph",
  "email": "okaforchekus@gmail.com"
}
```

## Running the Automated Tests

The application uses professional Go testing idioms to evaluate routing, headers, success schemas, and validation errors without spinning up an actual network port.

To execute the table-driven test suite, save the test code to `main_test.go` and run:

```bash
go test -v
```

