FROM golang:alpine as build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app/
COPY . .
RUN go get ./...
RUN go build -o ./bin/GoCSVServer ./cmd/CSVServer/main.go

FROM alpine:3.9
WORKDIR /go/bin
COPY --from=build /go/src/app/bin/GoCSVServer ./GoCSVServer
EXPOSE 80
ENTRYPOINT /go/bin/GoCSVServer
