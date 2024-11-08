FROM golang:latest as builder
COPY . /app
WORKDIR /app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o mockly /app/src/main.go

FROM alpine
ARG ENVIRONMENT
LABEL maintainer = "mockly"
WORKDIR /app
COPY --from=builder /app/mockly /app
ENTRYPOINT ./mockly