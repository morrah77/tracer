FROM golang:1.9.1
WORKDIR /tracer
ENV PATH=$PATH:/tracer/bin
ENV GOPATH=/go:/tracer
COPY ./ /tracer/
RUN go install tracer/services/...
RUN rm -rf src pkg
