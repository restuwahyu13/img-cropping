FROM golang:alpine

ENV GO111MODULE="on" \
    CGO_ENABLED="0"

COPY go.mod ./
COPY . ./

RUN apk update \
    && apk -u list \
    && apk upgrade \
    && apk --no-cache add build-base

RUN make build

EXPOSE 8080
ENTRYPOINT ["./main"]