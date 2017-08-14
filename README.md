# httpserver
Simple http server that serves the selected folder content

    go get github.com/0xC0D3/httpserver
    go build github.com/0xC0D3/httpserver

## Usage:
- **-port int** HTTP Server port. (default 8080)
- **-root string** HTTP Server root directory.

### Example:
    httpserver -root="%GOPATH%/src/github.com/0xC0D3/httpserver/test"
