FROM golang:1.23-alpine AS build

COPY ./server /app/server

WORKDIR /app/server/gateway
RUN go mod download
RUN go build -o gateway

FROM scratch

WORKDIR /
COPY --from=build /app/server/gateway /
EXPOSE 8080
ENTRYPOINT [ "./gateway" ]
