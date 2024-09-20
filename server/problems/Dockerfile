FROM golang:1.23-alpine AS build

COPY ./server /app/server
WORKDIR /app/server/problems
RUN go mod download
RUN go build -o problems

FROM scratch

WORKDIR /
COPY --from=build /app/server/problems /
EXPOSE 4000
ENTRYPOINT [ "./problems" ]