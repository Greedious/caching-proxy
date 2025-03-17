# Caching Reverse Proxy Server

A simple Caching Reverse Proxy server implemented in Go. This server acts as an intermediary between clients and backend servers, caching HTTP responses in an in-memory map to improve performance and reduce the load on backend services.

## Features

- Caches HTTP responses to reduce redundant requests to the backend server.
- Uses an in-memory map for fast data retrieval.
- Handles basic proxying of HTTP requests.
- Supports GET requests for caching responses.

## Installation

### Prerequisites

- [Go](https://golang.org/) installed (version 1.18+ recommended)

### Clone the Repository

```sh
git clone https://github.com/yourusername/caching-reverse-proxy.git
cd caching-reverse-proxy
```

### Build the Server

```sh
go build -o caching-proxy
```

## Usage

Start the caching proxy server with the following command:

```sh
caching-proxy --port <number> --origin <url>
```

### Example

```sh
caching-proxy --port 3000 --origin http://dummyjson.com
```

## Configuration

You can configure the server using command-line arguments:

- `--port <number>`: Specifies the port on which the proxy server runs.
- `--origin <url>`: Defines the backend server URL that the proxy will fetch data from.

## Example Usage

```sh
curl -X GET http://localhost:3000/products?limit=10&offset=0
```

If the response is not cached, the proxy fetches it from the backend. The next request for the same resource will return the cached response.

## Future Improvements

- Implement a cache eviction policy (e.g., LRU, TTL-based expiration).
- Add support for more HTTP methods.
- Introduce persistence using Redis or a file-based cache.
- Enhance error handling and logging.

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Feel free to open issues and submit pull requests.

## Author

Your Name - [Your GitHub Profile](https://github.com/Greedious)
