FROM golang:1.23-alpine as build

WORKDIR /
COPY go.mod ./
RUN go mod download
COPY . .
RUN go get -u
RUN go build -o problems

FROM scratch

WORKDIR /
COPY --from=build /problems /
EXPOSE 4000
ENTRYPOINT [ "./problems" ]