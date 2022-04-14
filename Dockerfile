FROM golang:1.18.1-bullseye AS build_step
ADD . /src
RUN cd /src && go build -o hn-to-rss

FROM alpine
WORKDIR /app
LABEL maintainer='matt@brokenintuition.com'
COPY --from=build_step /src/hn-to-rss /app
ENTRYPOINT ./hn-to-rss
