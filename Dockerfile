FROM golang:1.16.0-alpine AS go_build

RUN apk --update --no-cache add git build-base tzdata openssh

ENV TZ Asia/Tokyo

WORKDIR /himo-ingame

COPY go.mod go.sum /himo-ingame/
RUN go mod download

COPY . /himo-ingame
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  make

FROM alpine:3.12.3

RUN apk --update --no-cache add curl nmap mysql-client bash ca-certificates jq && \
  update-ca-certificates

WORKDIR /himo-ingame
COPY ./index.html ./index.html
COPY ./entry-point.sh ./entry-point.sh
COPY --from=go_build /himo-ingame/bin/ingame ./bin/

RUN chmod 755 ./entry-point.sh
ENTRYPOINT [ "./entry-point.sh" ]
