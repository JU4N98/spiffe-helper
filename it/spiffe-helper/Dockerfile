FROM golang:1.17-alpine AS spiffe-helper

# Set the desired version of spiffe-helper
ENV SPIFFE_HELPER_VERSION="2.1.6"

# Copy the spiffe-helper source code to the container
COPY ./ /service/

WORKDIR /service

# Build the spiffe-helper binary with the specified version
RUN go build -tags netgo -a -v -o /service/spiffe-helper \
    --ldflags "-X main.version=${SPIFFE_HELPER_VERSION}" \
    ./cmd/spiffe-helper

