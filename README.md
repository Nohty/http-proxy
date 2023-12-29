# HTTP Proxy Server in Go

This repository contains the source code for a simple HTTP proxy server written in Go. The server is designed to forward HTTP requests to specified URLs and return the responses. It supports various HTTP methods including GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS.

## Features

- **Multiple HTTP Methods:** Handles GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS requests.
- **Dynamic URL Routing:** Forwards requests to URLs specified in query parameters.
- **Header Forwarding**: Copies headers from the original request to the forwarded request and vice versa.
- **Cross-Origin Resource Sharing (CORS):** Includes CORS headers in responses to support cross-origin requests in web applications.
- **Error Handling:** Provides informative error messages for various failure scenarios.

## Requirements

- Go (version 1.21 or later)

## Installation

1. **Clone the Repository:**

```bash
git clone https://github.com/yourusername/your-repo-name.git
```

2. **Navigate to the Directory:**

```bash
cd your-repo-name
```

## Usage

To start the server, run:

```bash
go run main.go
```

You can also specify a custom port using the -p or --port flag:

```bash
go run main.go -p 8081
```

## API Endpoints

The server listens for requests on the root path `/` and forwards them based on the HTTP method and the `url` query parameter.

### Example Requests

#### GET Request:

```bash
http://localhost:8080/?url=https://example.com
```

#### POST Request with Data:

```bash
http://localhost:8080/?url=https://example.com/api/data
body: {"name": "John Doe", "age": 25}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Note:** This is a simple project intended for learning purposes. It might not be suitable for production use without additional security and performance considerations.
