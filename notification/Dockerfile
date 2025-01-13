# Stage 1: Build Stage
FROM golang:1.23-alpine AS builder

COPY shared /go/src/github.com/azeezlala/assessment/shared
COPY notification /go/src/github.com/azeezlala/assessment/notification

WORKDIR /go/src/github.com/azeezlala/assessment/notification

RUN go mod tidy
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates libc6-compat tzdata

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/azeezlala/assessment/notification/main .

EXPOSE 20021

CMD [ "./main" ]
