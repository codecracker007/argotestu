FROM golang:1.21

WORKDIR /usr/src/app

COPY mainu.go .

RUN go build mainu.go

EXPOSE 8081

CMD ["./mainu"]