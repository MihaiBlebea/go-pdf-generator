
# Build container
FROM golang:1.16.2-buster AS build_base

RUN apt-get install git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/diploma .

# Start fresh from a smaller image for the runtime container
FROM debian:buster

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    ca-certificates \
    xvfb \
    libfontconfig \
    wkhtmltopdf

RUN update-ca-certificates

WORKDIR /app

VOLUME [ "/storage" ]

USER nobody

COPY --from=build_base --chown=nobody /tmp/app/out/diploma /app/diploma

ENTRYPOINT ["./diploma", "generate", "-c=10000"]