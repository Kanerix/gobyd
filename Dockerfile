ARG SERVICE_NAME

FROM golang:1.23.3-alpine3.20 AS builder
WORKDIR /build

RUN apk update && apk upgrade --no-cache && \
    apk add --no-cache make

WORKDIR /build

ADD . .

ENV SERVICE_NAME=${SERVICE_NAME}
RUN if [ -z $SERVICE_NAME ]; then \
        echo "Missing $SERVICE_NAME build-arg"; exit 1; \
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

ENV SERVICE_NAME=${SERVICE_NAME}
ENV SERVICE_ADDR=0.0.0.0:8080
EXPOSE 8080

CMD ["bin"]
