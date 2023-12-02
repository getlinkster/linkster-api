FROM golang:1.21.3

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o linkster

CMD ["./linkster"]