# 🚀 Simple Go JSON API Server

Hey there! This is a minimal, lightweight HTTP server written in Go that serves data in JSON format instead of traditional HTML. It is a perfect starting point if you are looking to build web APIs, microservices, or backends for frontend frameworks like React, Vue, or mobile apps.

---

## Features

* **Zero Third-Party Dependencies**: Built entirely using Go's standard library (`net/http` and `encoding/json`).
* **Proper Headers**: Automatically tells clients to expect JSON data by explicitly setting the `application/json` Content-Type.
* **Streamlined Encoding**: Uses `json.NewEncoder` to stream data directly into the response writer, keeping memory usage low.

---

##  Project Structure

The project maps data into a native Go struct before converting it to raw JSON:

* **`UserResponse`**: A custom Go structure specifying the shape of our data. 
* **`json:"message"`**: Struct tags ensure our fields stay lowercase in the final JSON output, matching standard API conventions.

---

## How to Run It

### 1. Run the server
Make sure you have Go installed on your machine. Save the code into a file named `main.go` and run the following command in your terminal:

```bash
go run main.go
```

You should see this confirmation message in your terminal:
> `Starting the server on port :8080`

### 2. Test the API response
While the server is running, open a separate terminal window and hit the server using `curl`:

```bash
curl http://localhost:8080
```

**Expected JSON Output:**
```json
{
  "message": "Welcome to updated http server in go with JSON,this is JSON data!",
  "status": 200
}
```

Alternatively, you can just open your browser and navigate to `http://localhost:8080` to see the text on your screen.

---

##  What's Happening Under the Hood?

1. **Routing**: `http.HandleFunc("/", HomePage)` listens for any incoming web requests on the root path.
2. **Setting Headers**: `w.Header().Set(...)` stops the browser from guessing the file type and forces it to parse the data purely as JSON.
3. **Encoding Data**: Instead of creating a temporary string, `json.NewEncoder(w).Encode(data)` converts the Go struct directly into bytes and sends it across the network immediately for maximum performance.
