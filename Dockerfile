FROM golang:1.13-alpine AS build

#Install git
RUN apk add --no-cache git
RUN go get github.com/photosojourn/go-json-microserve-image
WORKDIR /go/src/github.com/photosojourn/go-json-microserve-image
RUN go build -o /bin/go-json-microservice

FROM golang:1.13-alpine
COPY --from=build /bin/go-json-microservice /bin/go-json-microservice
EXPOSE 8080
ENTRYPOINT ["/bin/go-json-microservice"]