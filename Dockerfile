FROM golang:1.25.4-alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

# Keep module cache efficient (optional for dev)
COPY go.mod go.sum ./
RUN go mod download || true

COPY . .

EXPOSE 8080
CMD ["go", "run", "./cmd/web"]
