# Build Stage
FROM golang:1.20 AS build

# Set the working directory
WORKDIR /app

# Copy all files to the working directory
COPY . .

# Download all dependencies
RUN go mod download

# Generate the query files
RUN go run cmd/gen/gen.go

# Set CGO_ENABLED to 1
ENV CGO_ENABLED=1

# Build the Go application
RUN go build -o /app/app ./cmd/app/app.go

# Runtime Stage
FROM busybox AS runtime

# Copy the binary from the build stage to the runtime stage
COPY --from=build /app/app /app

# Copy the configuration files from the build stage to the runtime stage
COPY --from=build /app/rbac/model.conf /rbac/model.conf
COPY --from=build /app/rbac/policy.csv /rbac/policy.csv

# Set the entry point to execute the binary directly
ENTRYPOINT ["/app"]
