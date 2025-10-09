# Stage 1: Build the Go binary
FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Ensure the binary is statically linked
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
# Debug step to list the files
RUN ls -l /app
# Stage 2: Create the final image using distroless
FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/main .
EXPOSE 80
CMD ["./main"]
