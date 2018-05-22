FROM golang:1.10-alpine AS builder

# Enable edge repository
COPY api/alpine.repo .
RUN cat alpine.repo >> /etc/apk/repositories

# Add dependencies
RUN apk add --no-cache \
  build-base \
  dep \
  git

WORKDIR /go/src/github.com/joaofnfernandes/autoredeploy
COPY api api
COPY pkg pkg

# Get dependencies, and build
# All dependencies should be vendored in. This is only needed if we
# don't include the pkg in the container image
# RUN dep ensure
RUN go build -v -a -o /usr/local/bin/api ./api

FROM alpine:latest
COPY --from=builder /usr/local/bin/api /usr/local/bin/api
# Default port
EXPOSE 8000
ENTRYPOINT ["/usr/local/bin/api"]
