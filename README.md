# init_golang

## Golang version

- Should use `1.18`

## Quick start with server side

- Step 1: Install necessary tools:

```
go mod tidy
```

- Step 2: Start the docker up:

```
make docker/up
```

- Step 3: Start server side:


```
With gin framework: make run/gin

With echo framework: make run/echo
```
