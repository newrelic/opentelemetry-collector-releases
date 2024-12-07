FROM golang:1.23 AS build
WORKDIR $GOPATH/main
COPY . .
RUN go env -w GOPROXY=direct
RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux go build -o=/bin/main .

FROM scratch
COPY --from=build /bin/main /bin/main

ENTRYPOINT ["/bin/main"]
