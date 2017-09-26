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

## Changelog:

### Alpha 0.0.2
TLS mode implemented, now our server can serve content securely.

#### CLI
- Supports `-tls` parameter to enable _**HTTP/TLS**_ mode, next parameters are required then:
    - `-cert="<path to the public cert.pem file>"`
    - `-pkey="<path to the private key.pem file>"`
    
#### Repo
- Repo name changed to _**webserver**_
- Branch for development version.

#### Code
- Added current version to [`main.go`](https://github.com/0xC0D3/webserver/blob/master/main.go#L3) in comments.

### Alpha 0.0.1
First release.
#### CLI
- Supports `-port=<http port number>` parameter (default 8080).
- Supports `-root="<web root directory path>"` parameter (Where index.html lives, no default value).