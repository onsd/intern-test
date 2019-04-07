FROM golang:1.12.1 as builder

ADD src /go/src/workdir/src
ENV GO111MODULE=on
ENV PORT=8080

WORKDIR /go/src/workdir/src

RUN CGO_ENABLED=0 go build -o APIServer

FROM alpine 
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/workdir/src/APIServer /APIServer
EXPOSE 8080
ENTRYPOINT [ "/APIServer" ]