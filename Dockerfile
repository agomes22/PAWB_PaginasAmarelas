FROM golang:1.11.1

WORKDIR /go/src/
COPY . .

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
