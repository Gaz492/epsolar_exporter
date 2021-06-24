FROM golang:1.16-alpine

RUN apk update \
    && apk add --no-cache wget curl git

WORKDIR /usr/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 2122
CMD ["epsolar_exporter"]