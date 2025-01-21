FROM alpine:3.18 as certs
RUN apk --update add ca-certificates

FROM alpine:3.18

RUN apk add --no-cache bash

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --chmod=755 nr-otel-collector /nr-otel-collector
COPY configs/nr-otel-collector-agent-linux.yaml /etc/nr-otel-collector/config.yaml
ENV EXTRA_COLLECTOR_ARGS ""
ENTRYPOINT ["/bin/sh", "-c", "exec /nr-otel-collector --config /etc/nr-otel-collector/config.yaml --feature-gates=-pkg.translator.prometheus.NormalizeName ${EXTRA_COLLECTOR_ARGS}"]
EXPOSE 4317 55678 55679
