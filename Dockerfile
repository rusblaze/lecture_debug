FROM golang:bullseye AS build-env
LABEL authors="aivanov0186"

RUN go install github.com/go-delve/delve/cmd/dlv@latest

ADD . /dockerdev
WORKDIR /dockerdev

RUN cd cmd && go build -gcflags="all=-N -l" -o /server

# Final stage
FROM debian:buster
EXPOSE 3333 40000
WORKDIR /
COPY --from=build-env /go/bin/dlv /
COPY --from=build-env /server /
COPY ./data /data
CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/server"]