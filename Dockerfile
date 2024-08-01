FROM golang:latest

RUN mkdir /build
WORKDIR /build

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 8000

CMD [ "./main" ]