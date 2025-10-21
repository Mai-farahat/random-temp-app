# Step 1: Use Go base image
FROM golang:1.25-alpine

# Step 2: Set working directory
WORKDIR /app

# Step 3: Copy go.mod and go.sum first
COPY go.mod go.sum* ./

# Step 4: Download dependencies (safe even if none)
RUN go mod download

# Step 5: Copy the rest of the code
COPY . .

# Step 6: Build the binary
RUN go build -o main .

# Step 7: Run
CMD ["./main"]
