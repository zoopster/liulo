# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# File Author / Maintainer
MAINTAINER tieubao

RUN go get github.com/tools/godep
# ENV GOPATH=`godep path`:$GOPATH

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/dwarvesf/liulo

# RUN go get github.com/astaxie/beego
# RUN go get github.com/beego/bee

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN godep restore
RUN bee run

# Run the outyet command by default when the container starts.
# ENTRYPOINT /go/bin/outyet

# Document that the service listens on port 8080.
EXPOSE 6060