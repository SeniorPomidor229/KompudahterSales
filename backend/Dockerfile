FROM golang:1.20.1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /comp

EXPOSE 8000

CMD [ "/comp" ]