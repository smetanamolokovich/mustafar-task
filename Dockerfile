FROM golang:1.17.2-alpine3.14

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o ./api ./cmd/api/

CMD ["./api"]