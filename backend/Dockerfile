## Stage 1: Build the Go application
#FROM golang:alpine AS build-env
#
#WORKDIR /app
#COPY . .
#COPY config /app/config
#
## Build the Go application and output it as "binary"
#RUN go build -o /app/binary
#
## Stage 2: Create a minimal image to run the application
#FROM alpine:3.14
#
## Expose port 9999
#EXPOSE 9999
#
#WORKDIR /
#
## Copy the binary from the build stage into the final image
#COPY --from=build-env /app/binary /app/binary
#COPY --from=build-env /app/config.yaml /app/config.yaml
#
## Define the command to run your application
#CMD ["/app/binary"]

# Menggunakan gambar Golang sebagai dasar
FROM golang:1.21

# Set kerja direktori di dalam kontainer
WORKDIR /app

# Menyalin semua berkas aplikasi Golang ke dalam kontainer
COPY . .

# Build aplikasi Golang
RUN go build -o app

# Port yang akan digunakan oleh aplikasi
EXPOSE 9999

# Perintah untuk menjalankan aplikasi
CMD ["./app"]
