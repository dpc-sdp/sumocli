# Build the manager binary
FROM golang:latest as builder

WORKDIR ${GOPATH}/src/github.com/dpc-sdp/sumocli

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY . ./

# Build
ARG TARGETPLATFORM=linux/amd64
RUN CGO_ENABLED=0 \
    GOOS=${TARGETPLATFORM%%/*} \
    GOARCH=${TARGETPLATFORM#*/} \
    go build -a -o /bin/sumocli ./cmd/sumocli

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

LABEL org.opencontainers.image.authors="Single Digital Presence"
LABEL org.opencontainers.image.description="A CLI application that lets you manage/automate your Sumo Logic tenancy."
LABEL org.opencontainers.image.source="https://github.com/dpc-sdp/sumocli"

WORKDIR /
COPY --from=builder /bin/sumocli .
USER nonroot:nonroot

ENTRYPOINT ["/sumocli"]
CMD ["--help"]