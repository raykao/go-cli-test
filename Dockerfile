FROM golang:1.17.5-bullseye as dev


RUN go install github.com/spf13/cobra/cobra@latest

WORKDIR /root/projects