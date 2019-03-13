FROM golang:1.8

RUN mkdir /go/src/goservice
RUN chdir /go/src/goservice
WORKDIR /go/src/goservice
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["goservice"]
