FROM golang:bookworm

WORKDIR /goban

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o /app cmd/main.go

ENTRYPOINT [ "/app" ]
