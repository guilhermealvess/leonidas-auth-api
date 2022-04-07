FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN sh execute-tests.sh

RUN go build -o /server

EXPOSE 8000
EXPOSE 50052

CMD [ "/server" ]