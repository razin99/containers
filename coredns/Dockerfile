FROM scratch

USER 1010:1010

COPY coredns /coredns
WORKDIR /
EXPOSE 53 53/udp
ENTRYPOINT ["/coredns"]
