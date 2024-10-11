# Goserve
Goserve is a simple command-line tool written in Go that serves files from a specified directory over HTTP or HTTPS. It provides a user-friendly way to quickly set up a web server for local development, file sharing, or testing. This is similar to the http server in python.(`python -m http.server`)


## Features
- Serve files from a specified directory.
- Options for configuring the host and port.
- Support for SSL with certificate and key files.
- Verbose logging for easier debugging.


##  Installation
Make sure you have Go installed on your system. You can download it from golang.org.

To install Goserve, run:
```
go install github.com/vshalt/goserve@latest`
```
Or build the binary with:
```
make build
```
Make sure that $GOPATH/bin is in your PATH to run the goserve command.

## Usage

goserve [OPTIONS]

Command-Line Arguments

```
--dir (default: current working directory): The directory to serve files from.
--host (default: "127.0.0.1"): The host to bind the server to.
--port (default: 3333): The port to serve on.
--ssl (default: false): Enable SSL for the server.
--cert (required if SSL is enabled): Path to the SSL certificate file.
--key (required if SSL is enabled): Path to the SSL key file.
--verbose (default: false): Enable verbose logging for debugging.
```

## Example
- Serve files from the current directory on localhost at port 3333:
```
goserve
```
- Serve files from a specific directory with verbose logging:
```
goserve --dir /path/to/directory --verbose
```
- Serve files over HTTPS:

```
goserve -dir /path/to/directory -s -c /path/to/cert.pem -k /path/to/key.pem
```

- Serve files on a different host and port:
```
goserve -d /path/to/directory -h 0.0.0.0 -p 3000
```

## License
This project is licensed under the MIT License. See the LICENSE file for more information.
