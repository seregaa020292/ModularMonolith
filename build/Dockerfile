# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22
ARG ALPINE_VERSION=3.19

#-------------------------------------------------------------------------------
# STAGE: BASE
#-------------------------------------------------------------------------------
FROM golang:${GO_VERSION}-alpine AS base
ARG TARGET_CGO_ENABLED=0
ARG TARGET_OS=linux
ARG TARGET_ARCH=amd64
WORKDIR /app
RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add ca-certificates tzdata curl \
    && update-ca-certificates
RUN --mount=type=cache,target=${GOMODCACHE} \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

#-------------------------------------------------------------------------------
# STAGE: CLI
#-------------------------------------------------------------------------------
FROM golang:${GO_VERSION}-alpine AS cli
ARG UID=1000
ARG GID=$UID
ARG GOLANGCI_LINT_VERSION=1.59.0
ARG GOOSE_VERSION=3.20.0
ARG OAPI_CODEGEN_VERSION=2.3.0
ARG JET_VERSION=2.11.1
ARG WIRE_VERSION=0.6.0
WORKDIR /app
RUN apk --update add curl
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin v${GOLANGCI_LINT_VERSION}
RUN curl -sSfL https://raw.githubusercontent.com/pressly/goose/master/install.sh | GOOSE_INSTALL=${GOPATH} sh -s v${GOOSE_VERSION}
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v${OAPI_CODEGEN_VERSION}
RUN go install github.com/go-jet/jet/v2/cmd/jet@v${JET_VERSION}
RUN go install github.com/google/wire/cmd/wire@v${WIRE_VERSION}
RUN addgroup -g ${GID} app && adduser -u ${UID} -G app -s /bin/sh -D app
USER app

#-------------------------------------------------------------------------------
# STAGE: TEST
#-------------------------------------------------------------------------------
FROM base AS test
RUN --mount=type=cache,target=${GOMODCACHE} \
    --mount=type=bind,target=. \
    go test -v -coverprofile=/tmp/coverage.txt ./... > /tmp/result.txt; \
    [[ $? -eq 0 ]] || { cat /tmp/result.txt; exit 1; }

FROM scratch AS test-export
COPY --from=test /tmp/coverage.txt /
COPY --from=test /tmp/result.txt /

#-------------------------------------------------------------------------------
# STAGE: BUILD
#-------------------------------------------------------------------------------
FROM base AS build
ENV CGO_ENABLED=$TARGET_CGO_ENABLED \
    GOOS=$TARGET_OS \
    GOARCH=$TARGET_ARCH
RUN --mount=type=cache,target=${GOMODCACHE} \
    --mount=type=cache,target=${GOCACHE} \
    --mount=type=bind,target=. \
    go build -o /bin/app ./cmd/app

#-------------------------------------------------------------------------------
# STAGE: DEVELOPMENT FROM ENV
#-------------------------------------------------------------------------------
FROM base AS development
ARG AIR_VERSION=1.52.2
RUN --mount=type=cache,target=${GOMODCACHE} \
    go install github.com/air-verse/air@v${AIR_VERSION}
ENV CGO_ENABLED=$TARGET_CGO_ENABLED \
    GOOS=$TARGET_OS \
    GOARCH=$TARGET_ARCH
CMD ["air", "-c", ".air.toml"]

#-------------------------------------------------------------------------------
# STAGE: PRODUCTION FROM ENV
#-------------------------------------------------------------------------------
FROM alpine:${ALPINE_VERSION} AS production
COPY --from=build /bin/app /bin/app
HEALTHCHECK --interval=30s --timeout=10s --retries=3 \
    CMD wget -qO- http://localhost:8080/health || exit 1
ENTRYPOINT ["/bin/app"]
