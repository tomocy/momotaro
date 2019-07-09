# Momotaro

a CLI client for [Kibidango](https://github.com/tomocy/kibidango) (a container from scratch)

### Environment
```
docker run -it -v $GOPATH:/go -w /go/src/github.com/tomocy/momotaro --rm --privileged golang:alpine
```

### Usage
- list all kibidangos
    ```
    go run main.go list
    ```
- create a kibidango
    ```
    go run main.go create ${container id}
    ```
- start a kibidango
    ```
    go run main.go start ${container id}
- kill a kibidango of given id with given signal
    ```
    go run main.go kill ${container id} ${signal}
    ```
- delete a kibidango
    ```
    go run main.go delete ${container id}
    ```