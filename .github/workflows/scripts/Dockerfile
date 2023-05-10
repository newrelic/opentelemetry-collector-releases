FROM debian:bullseye

RUN apt-get update \
    && apt-get -y install \
        rpm \
        gnupg2 \
        gpg-agent \
        debsigs

RUN mkdir -p /srv/scripts

COPY .github/workflows/scripts/sign.sh /srv/scripts/sign.sh
RUN chmod +x /srv/scripts/sign.sh
