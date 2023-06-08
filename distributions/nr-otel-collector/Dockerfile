FROM alpine:3.18 as certs
RUN apk --update add ca-certificates

FROM alpine:3.18

RUN apk add --no-cache bash

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --chmod=755 nr-otel-collector /nr-otel-collector
COPY configs/nr-otel-collector-agent-linux.yaml /etc/nr-otel-collector/config.yaml
ENTRYPOINT ["/nr-otel-collector"]
CMD ["--config", "/etc/nr-otel-collector/config.yaml", "--feature-gates=-pkg.translator.prometheus.NormalizeName"]
EXPOSE 4317 55678 55679
