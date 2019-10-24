# go application
FROM golang:1.12.0-alpine3.9
RUN mkdir /teamCMP
ADD . /teamCMP
WORKDIR /teamCMP
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get gopkg.in/yaml.v2
RUN go get github.com/Ahrimal/teamCMP
RUN go build .
ENTRYPOINT ["/teamCMP/teamCMP"]