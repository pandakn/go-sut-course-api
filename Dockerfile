FROM golang:alpine as builder
WORKDIR /go/src
COPY . .
RUN go get \ 
    &&  go build -o /go/bin/app

FROM alpine
COPY --from=builder /go/bin/app /web/app
COPY .env.prod /web

EXPOSE 8080

ENTRYPOINT [ "/web/app", "/web/.env.prod" ]