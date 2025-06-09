FROM golang:1.24.4-alpine3.22
WORKDIR /app
COPY /appForDocker/endlessHello.go /app/helloDocker.go
ENTRYPOINT [ "go", "run", "helloDocker.go"]