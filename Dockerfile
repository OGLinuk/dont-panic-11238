FROM golang:1.15
ADD . /go/src/dont-panic-11238
WORKDIR /go/src/dont-panic-11238
RUN go build -o dont-panic-11238
CMD ["./dont-panic-11238& && docker-compose up --build --remove-orphans"]