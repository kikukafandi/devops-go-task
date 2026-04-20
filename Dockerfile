# ==========================================
# Stage 1: Build the Go binary
# ==========================================
FROM golang:1.22-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.sum (jika ada) untuk caching dependencies
COPY go.mod ./
# RUN go mod download # (Uncomment baris ini jika kamu memakai external package)

# Copy seluruh source code
COPY . .

# Build aplikasi dengan CGO_ENABLED=0 agar menghasilkan static binary yang bisa jalan di scratch/alpine murni
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# ==========================================
# Stage 2: Create the final lightweight image
# ==========================================
FROM alpine:latest

WORKDIR /app

# Copy binary hasil build dari Stage 1
COPY --from=builder /app/main .

# Expose port yang digunakan aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]