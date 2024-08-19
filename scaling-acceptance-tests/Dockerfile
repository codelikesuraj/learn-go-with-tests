FROM golang:1.22-alpine

ARG bin_to_build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/${bin_to_build}/*.go

EXPOSE 8080

CMD [ "./svr" ]