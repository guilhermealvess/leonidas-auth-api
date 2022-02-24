FROM golang:1.17

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN . execute-tests.sh
RUN go build -o /usr/local/bin/app

COPY .env /usr/local/bin/.env

EXPOSE 50052

CMD ["app"]