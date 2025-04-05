FROM golang:1.24.2

LABEL org.opencontainers.image.documentation="https://github.com/itning/dify-workflow-trigger/blob/main/README.md"
LABEL org.opencontainers.image.authors="itning"
LABEL org.opencontainers.image.source="https://github.com/itning/dify-workflow-trigger"
LABEL org.opencontainers.image.title="dify-workflow-trigger"
LABEL org.opencontainers.image.description="Timed tasks trigger Dify workflow execution"
LABEL org.opencontainers.image.licenses="Apache License 2.0"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/dify-workflow-trigger ./...

CMD ["dify-workflow-trigger", "--config", "/app/config.json"]