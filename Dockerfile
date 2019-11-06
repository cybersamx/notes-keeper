FROM golang:1.13.4-alpine as build

# Copy the source files
WORKDIR /notes-keeper
COPY . ./

# Install dependencies
RUN go mod download

# Build
WORKDIR /notes-keeper/app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../bin/notes-keeper

FROM alpine:latest

# Run as non-privileged user
RUN mkdir -p /notes-keeper \
    && adduser -D gouser \
    && chown -R gouser:gouser /notes-keeper
USER gouser

WORKDIR /notes-keeper/html
COPY --from=build --chown=gouser /notes-keeper/html ./

# Copy the built binary from previous stage
WORKDIR /notes-keeper/bin/
COPY --from=build --chown=gouser /notes-keeper/bin/notes-keeper ./

# Run
EXPOSE 8000
CMD ["./notes-keeper"]