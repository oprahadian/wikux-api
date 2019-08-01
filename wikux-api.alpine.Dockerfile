FROM golang:alpine

WORKDIR /go/src/wikux-api

COPY . .

RUN apk add dep 
RUN dep ensure
RUN go install .

CMD ["wikux-api"]