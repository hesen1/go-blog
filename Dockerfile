FROM golang:1.12.9

WORKDIR /blog
COPY . /blog
ENV GOPROXY https://goproxy.cn

# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build .

EXPOSE 8000

CMD ["./blog"]