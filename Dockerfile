FROM alpine:3.17

WORKDIR /app
RUN apk add go && \
    apk add make git

COPY . /app
RUN make build

EXPOSE 9000
ENTRYPOINT ["./bin/music_service"]