FROM golang:1.13-alpine3.10 AS build

RUN apk add --no-cache git openssh

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app/zendesk

COPY . .

RUN go mod download
RUN go build -o bin/app main.go

FROM alpine:3.10

RUN adduser -D -u 1000 zendesk

WORKDIR /home/zendesk

USER zendesk

COPY --from=build /app/zendesk/bin/app .
COPY --from=build /app/zendesk/data ./data

CMD [ "/home/zendesk/app" ]
