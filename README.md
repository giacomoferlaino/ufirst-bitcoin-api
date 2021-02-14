# ufirst-bitcoin-api

## How to run

The server can be run in 2 ways:

- using a Go compiler
- using Docker

**The default port for the server is: `8080`**

### Using go compiler

To run the server using the Go compiler, use the following command:

```bash
# go run . <port number>
go run . 8080
```

### Using Docker

To run the server using Docker, using the **docker-compose** command:

```bash
# to define a custom port edit the PORT variable inside the  .env file
docker-compose up --build
```

using the **docker** command:

```bash
# docker build -t ufirst/bitcoin-api . && docker run -it -p <port_number>:<port_number> --env PORT=<port_number> ufirst/bitcoin-api
docker build -t ufirst/bitcoin-api . && docker run -it -p 8080:8080 --env PORT=8080 ufirst/bitcoin-api
```
