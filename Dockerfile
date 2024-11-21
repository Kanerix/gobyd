FROM golang:1.23.3-alpine3.20 AS builder
ARG SERVICE
WORKDIR /build

RUN apk update && apk upgrade --no-cache && \
    apk add --no-cache make

WORKDIR /build

ADD . .

ENV SERVICE=${SERVICE}
RUN if [ -z $SERVICE ]; then \
        echo "Missing $SERVICE build-arg"; exit 1; \
    fi

RUN make docker-build


FROM alpine:3.20 AS runner
WORKDIR /var/app

COPY --from=builder /build/bin .

RUN mkdir logs

RUN addgroup -S app-group && \
    adduser -S app-user -G app-group && \
    chown -R app-user:app-group /var/app

USER app-user

EXPOSE 8080

ENV GIN_MODE=release

CMD ["bin"]
