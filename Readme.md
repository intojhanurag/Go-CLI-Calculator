# Go Calculator 🧮

![Go CI](https://github.com/intojhanurag/Go-CLI-Calculator/actions/workflows/go.yml/badge.svg)

A simple CLI calculator built with Go — supports +, -, *, /.  
Dockerized with a multi-stage build and tested using GitHub Actions.

...
## Run

```bash
go run cmd/main.go
```

Or with Docker:

```bash
docker build -t calculator .
docker run -it calculator
```

## Test
```bash
go test ./...
```

...