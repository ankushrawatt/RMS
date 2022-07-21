FROM golang:1.16.6-alpine3.14
WORKDIR /rms

COPY go.mod ./
COPY go.sum ./

RUN  go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./rms

CMD echo 'hii, i am running.'
EXPOSE 8080
ENTRYPOINT ["./rms"]

