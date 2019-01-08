FROM golang:1.11.1 AS build
WORKDIR /mnt
COPY . .
RUN CGO_ENABLED=0 make build.all

FROM alpine:3.7
WORKDIR /opt
RUN apk add --no-cache ca-certificates
COPY --from=build /mnt/bin/* /usr/bin/
CMD ["server"]
