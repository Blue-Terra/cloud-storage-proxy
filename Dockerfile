FROM golang:1.16-alpine

WORKDIR /app 

COPY ./src /app/src
COPY go.mod ./
COPY go.sum ./

RUN go mod download 

RUN go build ./src/main.go 

EXPOSE 8080 

CMD [ "./main" ]