FROM golang:1.19-alpine as builder

RUN go version
ENV GOPATH=/

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o auth-app ./cmd/auth

EXPOSE 8008

CMD ["./auth-app"]