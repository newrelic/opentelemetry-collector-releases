FROM alpine:3.18 as certs
RUN apk --update add ca-certificates

FROM alpine:3.18

RUN apk add --no-cache bash

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --chmod=755 nrdot-collector-host /nrdot-collector-host
COPY config.yaml /etc/nrdot-collector-host/config.yaml
ENTRYPOINT ["/nrdot-collector-host"]
CMD ["--config", "/etc/nrdot-collector-host/config.yaml"]
# `4137` and `4318`: OTLP
EXPOSE 4317 4318
