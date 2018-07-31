FROM golang:latest

MAINTAINER Martin Tsai "cauu@163.com"

RUN apt-get update && apt-get install -y git

RUN cd $HOME && git clone 