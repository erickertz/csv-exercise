FROM golang:1.13.5 as builder

# Force the go compiler to use modules
ENV GO111MODULE=on

# Set the working directory outside $GOPATH to enable the support for modules
WORKDIR /src/scoircsvjson/

# Copy source code
COPY . .

# Install dependencies
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

# Unit test
RUN CGO_ENABLED=1 go test -cover -race ./...

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o scoircsvjson ./app/cmd/main.go


# Build a smaller container with s6-overlay
FROM project42/s6-alpine:3.8

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /src/scoircsvjson/scoircsvjson .

CMD ["./scoircsvjson"]
