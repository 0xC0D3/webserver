# httpserver
_**version:** Alpha 0.0.2_

Simple http server that serves the selected folder content

    go get github.com/0xC0D3/httpserver
    go build github.com/0xC0D3/httpserver

## Usage:
- **-cert _string_** TLS Public Cert file path.
- **-pkey _string_** TLS Private Key file path.
- **-port _int_** HTTP/HTTPS Server port. (default 8080)
- **-root _string_** Content root directory.
- **-tls** TLS Enable.

When `-tls` is set, you must specified `-cert` & `-pkey` params. 

### Example:
    httpserver -root="%GOPATH%/src/github.com/0xC0D3/httpserver/test" -port=80
