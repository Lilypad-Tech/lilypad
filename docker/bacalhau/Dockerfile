FROM docker:dind AS base

RUN apk update
RUN apk add wget
RUN apk add bash

ADD https://github.com/bacalhau-project/bacalhau/releases/download/v1.3.2/bacalhau_v1.3.2_linux_amd64.tar.gz .
RUN tar xfv bacalhau_v1.3.2_linux_amd64.tar.gz
RUN mv bacalhau /usr/local/bin

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD wget http://localhost:1234/api/v1/agent/alive -q || exit 1

ENTRYPOINT [ "bacalhau" ]