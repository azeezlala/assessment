# Stage 1: Build Stage
FROM golang:1.23-alpine AS builder

# Copy the source code for 'shared' and 'api'
COPY shared /go/src/github.com/azeezlala/assessment/shared
COPY api /go/src/github.com/azeezlala/assessment/api

WORKDIR /go/src/github.com/azeezlala/assessment/api

RUN go mod tidy
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates libc6-compat tzdata

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/azeezlala/assessment/api/main .

EXPOSE 8082

CMD [ "./main" ]