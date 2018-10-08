# STEP 1 build executable binary
FROM golang:alpine as builder
COPY . $GOPATH/src/dummy/
WORKDIR $GOPATH/src/dummy/
#build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/dummy
# STEP 2 build a small image
# start from scratch
FROM scratch
# Copy our static executable
COPY --from=builder /go/bin/dummy /go/bin/dummy
EXPOSE 9090
ENTRYPOINT ["/go/bin/dummy"]