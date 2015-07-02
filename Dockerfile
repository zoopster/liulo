# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# File Author / Maintainer
MAINTAINER tieubao

RUN go get github.com/tools/godep
# ENV GOPATH=`godep path`:$GOPATH

# Copy the local package files to the container's workspace.
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee

WORKDIR /go/src/github.com/dwarvesf/liulo
ADD . /go/src/github.com/dwarvesf/liulo
# RUN godep restore

# RUN bee run
# ENTRYPOINT bee run

# Document that the service listens on port 8080.
EXPOSE 8080
EXPOSE 8088

CMD bee run