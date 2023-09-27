# A MULTI-STAGE DOCKERFILE

# STAGE 1
# use alpine image as builder
FROM golang:alpine AS builder

# golang specific variables
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# current working directory is /build in the container
WORKDIR /build

# copy over go.mod and go.sum (module dependencies and checksum)
# over to working directory
COPY go.mod .
COPY go.sum .

# download the dependencies
RUN go mod download

# copy our application code into the container
COPY . .

# building the binary called "main"
RUN go build -o main .


# STAGE 2
# Build a small image
FROM scratch

# copy from stage-1 image

COPY --from=builder /build/config.yaml .
COPY --from=builder /build/main /

# expose the port to run the application on
EXPOSE 4000

# Command to run
ENTRYPOINT ["/main"]