FROM golang:latest

RUN apt-get update && apt-get install -y wget
RUN wget https://dist.ipfs.io/go-ipfs/v0.10.0/go-ipfs_v0.10.0_linux-amd64.tar.gz
RUN tar xvfz go-ipfs_v0.10.0_linux-amd64.tar.gz
RUN mv go-ipfs/ipfs /usr/local/bin/ipfs
RUN rm -rf go-ipfs_v0.10.0_linux-amd64.tar.gz go-ipfs

WORKDIR /app
COPY . .

ENV FILENAME=""
ENV API_PORT=""
ENV GATEWAY_PORT=""


RUN go build -o daemonBuild


CMD ["./daemonBuild"]


