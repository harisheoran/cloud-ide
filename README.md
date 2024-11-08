## Setup and run locally.
### Requirements
- Go version *go1.23.2*
- Install *air* package for live reload
    ```
    go install github.com/air-verse/air@latest
    ```

- Clone the repository.

1. Run the *server*
```
cd cloud-ide
```
- Install the dependencies
```
go mod download
```
- Run the server
```
air --build.cmd "go build -o bin/ide main.go" --build.bin "./bin/ide"
```

- Install ***wscat*** cli utility for testing websocket connection
```
wscat -c ws://localhost:5000/shell
```