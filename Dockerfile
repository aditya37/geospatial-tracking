FROM golang:1.16.4

WORKDIR /geospatial-tracking
COPY . .
ENTRYPOINT ["./geospatial-tracking"]