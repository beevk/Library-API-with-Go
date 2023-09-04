# ---- Build Stage ----
FROM golang:1.21 AS build

WORKDIR /app

# Copy only the Go module files to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application with optimizations
RUN CGO_ENABLED=0 go build -o /app/main -ldflags="-s -w" .

# ---- Final Stage ----
FROM scratch

# Copy the built binary from the build stage into the final image
COPY --from=build /app/main /app/main

# Expose the port your Go application will run on (if needed)
EXPOSE 8080

# Define the command to run your Go application
CMD ["/app/main"]
