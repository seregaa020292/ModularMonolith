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
