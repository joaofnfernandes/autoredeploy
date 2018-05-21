FROM golang:1.10-alpine AS builder

# Enable edge repository
COPY alpine.repo .
RUN cat alpine.repo >> /etc/apk/repositories

# Add dependencies
RUN apk add --no-cache \
  build-base \
  dep \
  git

WORKDIR /go/src/github.com/api
COPY . .
# Get dependencies, and build
RUN dep ensure
RUN go build -v -a -o /usr/local/bin/api .

FROM alpine:latest
COPY --from=builder /usr/local/bin/api /usr/local/bin/api
# Default port
EXPOSE 8000
ENTRYPOINT ["/usr/local/bin/api"]
