FROM golang:1.11 as builder

WORKDIR /go/src/github.com/Scout24/cgroup-metrics-reporter
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install -a -ldflags '-extldflags "-s -w -static"' .

FROM busybox
COPY --from=builder /go/bin/cgroup-metrics-reporter /usr/local/bin/cgroup-metrics-reporter
COPY entrypoint.sh /usr/local/bin/

EXPOSE 9301

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
