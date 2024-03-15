FROM golang:1.22

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /geoIP
CMD ["/geoIP"]