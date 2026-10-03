FROM quay.io/coreos/etcd:v3.5.15 AS etcd-source
FROM alpine:3.20

RUN apk add --no-cache \
  bash \
  curl \
  jq \
  kubectl \
  yq

COPY --from=etcd-source /usr/local/bin/etcdctl /usr/local/bin/etcdctl

CMD ["/bin/bash"]
