FROM alpine:latest as base

RUN apk add --no-cache ca-certificates && apk upgrade --no-cache

WORKDIR /app

FROM base as build

RUN apk add --no-cache go musl-dev

ENV GOPATH /app

ADD http.go /app

RUN go build -o /app/http .

FROM base as release

COPY --from=build /app/http /app/http

RUN chown -R nobody:nogroup /app

USER nobody

CMD /app/http $SESSIONS $URL
