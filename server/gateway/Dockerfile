FROM golang:1.23-alpine as build

WORKDIR /
COPY go.mod .
RUN go mod download
COPY  . .
RUN go build -o gateway

FROM scratch

WORKDIR /
COPY --from=build /gateway /
EXPOSE 8080
ENTRYPOINT [ "./gateway" ]