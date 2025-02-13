FROM golang:1.23

COPY . /go/src/github.com/microservices-demo/payment/

RUN go install github.com/DataDog/orchestrion@latest

RUN cd /go/src/github.com/microservices-demo/payment/ \
    && orchestrion pin \
    # && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app github.com/microservices-demo/payment/cmd/paymentsvc
    && CGO_ENABLED=0 GOOS=linux go build -toolexec="orchestrion toolexec" -a -installsuffix cgo -o /app github.com/microservices-demo/payment/cmd/paymentsvc

CMD ["/app", "-port=8080"]

EXPOSE 8080