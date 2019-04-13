############################################################
# Dockerfile for Github Golang API
# Based on golang latest
############################################################

# Set the base image to Golang Latest
FROM golang:latest

# File Author / Maintainer
MAINTAINER joshua@hauptj.com


################## Environment Variables ###################

ENV TOKEN=${TOKEN}

################## Begin Installation ######################

# Install Go Dependencies
RUN go get github.com/gorilla/context \
  github.com/gorilla/mux \
  github.com/google/go-github/github \
  golang.org/x/oauth2 \
  golang.org/x/net/context/ctxhttp


# Copy Golang source code
RUN mkdir /githubAPI
ADD API /githubAPI/

WORKDIR /githubAPI

# build API
RUN go build -o main .

# Run API
CMD ["/githubAPI/main"]

HEALTHCHECK CMD curl --fail http://localhost:8880/followers/hauptj || exit 1
