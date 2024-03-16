FROM golang:1.22

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o /geoip
CMD ["/bin/sh","-c","sleep 1d"]