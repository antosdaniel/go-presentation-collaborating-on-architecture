# Build production binary
FROM golang:1.22 as builder

WORKDIR /src/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=${version}" -v -o ./bin/main .

# Production image
FROM gcr.io/distroless/base-debian11:nonroot
COPY --from=builder /src/bin/main /bin/main

CMD ["/bin/main"]