# Momotaro

a CLI client for [Kibidango](https://github.com/tomocy/kibidango) (a container from scratch)

### Environment
```
docker run --rm -it -v $GOPATH:/go -w /go/src/github.com/tomocy/momotaro --privileged golang:alpine
```

### Usage
- create a kibidango
    ```
    go run main.go create
    ```