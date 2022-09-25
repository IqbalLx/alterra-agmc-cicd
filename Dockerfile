FROM golang:1.18-alpine3.16 AS build

WORKDIR /app

COPY ./go.* .
RUN go mod download

COPY . .
RUN go build -o alterra-agmc

FROM alpine:3.16 AS prod

RUN apk update
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
RUN apk add --update bash

WORKDIR /app
COPY --from=build /app/alterra-agmc .

CMD [ "./alterra-agmc"]
