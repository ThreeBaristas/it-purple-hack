FROM golang:1.22.1-alpine AS build
WORKDIR /app
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/purple .

# HANDLER
FROM alpine:3.19 AS purple
COPY --from=build /app/purple /main
EXPOSE 3000
RUN mkdir users
ENTRYPOINT ["/main"]

