FROM golang:1.15-alpine AS GO_BUILD
COPY . /server
WORKDIR /server/server
RUN apk add git
RUN apk update
ENV GOBIN /go/bin
RUN go get -u github.com/chromedp/chromedp
RUN set -eux && \
    GOOS=linux GOARCH=amd64 go build -o /go/bin/server/server

FROM chromedp/headless-shell:latest
WORKDIR /app
COPY --from=GO_BUILD /go/bin/server/ ./
RUN apt-get update
RUN apt-get install dumb-init -y
RUN apt-get install -y gcc musl-dev make
RUN ln -s /usr/lib/x86_64-linux-musl/libc.so /lib/libc.musl-x86_64.so.1
ENTRYPOINT ["dumb-init", "--"]
CMD ["./server"]
EXPOSE 8080

#FROM chromedp/headless-shell:latest
#RUN apt update
#RUN apt  install -y git
#RUN apt install -y dumb-init
#RUN apt install -y golang-go -v 1.16
#COPY ./server/ ./
#RUN go build -o /
#ENTRYPOINT ["dumb-init", "--"]
#CMD ["./server"]
#EXPOSE 8080

#FROM golang:1.15-alpine AS GO_BUILD
#COPY . /server
#WORKDIR /server/server
#RUN set -eux && GOOS=linux go build -o /go/bin/server/server
#
#FROM chromedp/headless-shell:latest
## Install dumb-init or tini
#WORKDIR /app
#COPY --from=GO_BUILD /go/bin/server/ ./
#RUN apt update
#RUN apt install dumb-init
## or RUN apt install tini
#ENTRYPOINT ["dumb-init", "--"]
## or ENTRYPOINT ["tini", "--"]
#CMD ["./server"]
#EXPOSE 8080


#FROM golang:1.15-alpine AS GO_BUILD
#COPY . /server
#WORKDIR /server/server
#RUN go build -o /go/bin/server/server
#
#FROM alpine:3.10
#WORKDIR app
#COPY --from=GO_BUILD /go/bin/server/ ./
#EXPOSE 8080
#CMD ./server