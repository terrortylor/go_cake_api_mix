FROM golang:latest

EXPOSE 8000
WORKDIR /usr/local/bin
COPY go_cake_api_mix go_cake_api_mix

ENTRYPOINT ./go_cake_api_mix
