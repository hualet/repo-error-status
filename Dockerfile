FROM debian:jessie-backports

COPY bin/* /usr/bin/

ENTRYPOINT ["/usr/bin/repo-error-status"]
