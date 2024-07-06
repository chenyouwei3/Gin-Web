FROM golang:1.20-alpine

WORKDIR "/go/src/app"

COPY . .
RUN go get -d -v ./..
RUN go install -v ./..

CMD["loopy-manager"]
ENTRYPOINT ["./app"]