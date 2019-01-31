From golang:alpine as builder

MAINTAINER Shrivatsa Upadhye "ishrivatsa@gmail.com"

# Future proofing by installing ca-cert for https support
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
COPY . $GOPATH/src/github.com/vmwarecloudadvocacy/catalogsvc
WORKDIR $GOPATH/src/github.com/vmwarecloudadvocacy/catalogsvc
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go build -o catalog .

FROM alpine
RUN mkdir app
RUN mkdir app/images
#Copy the executable uilt from the previous image
COPY --from=builder /go/src/github.com/vmwarecloudadvocacy/catalogsvc/catalog /app
COPY --from=builder /go/src/github.com/vmwarecloudadvocacy/catalogsvc/images /app/images
WORKDIR /app
EXPOSE 80
EXPOSE 8082
CMD ["./catalog"]