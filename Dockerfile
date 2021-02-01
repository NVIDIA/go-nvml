ARG GOLANG_VERSION=0.0.0
FROM golang:${GOLANG_VERSION}

# Get the supported version of c-for-go. Here we force the use of `GO111MODULE` for go get
# to support the @VERSION syntax.
ARG C_FOR_GO_TAG=master
RUN GO111MODULE=on go get github.com/xlab/c-for-go@${C_FOR_GO_TAG}
# Set the permissions on the go module path to ensure that this is accessible from
# our user containers.
RUN chmod -R a+rx /go/pkg/mod
