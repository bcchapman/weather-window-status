# syntax=docker/dockerfile:1
FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN go build -o /window-monitor

# Run
# CMD ["bash"]
ENTRYPOINT [ "tail", "-f", "/dev/null" ]