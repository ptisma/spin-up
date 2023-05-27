FROM golang:1.20.3-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN go build -o api-server ./cmd/api-server/

EXPOSE 8080 



CMD [ "./api-server" ]