ARG expose_via=local

FROM golang:latest AS base
WORKDIR /usr/src/app

ARG arch=amd64
ARG network=dev

COPY . .
RUN go mod download && go mod verify
RUN go build -v .
RUN go install

RUN (curl -Ls --tlsv1.2 --proto "=https" --retry 3 https://cli.doppler.com/install.sh || wget -t 3 -qO- https://cli.doppler.com/install.sh) | sh

RUN touch run
RUN echo "#!/bin/bash" >> run

FROM base AS expose-cloudflare
RUN curl -L --output cloudflared.deb https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-${arch}.deb
RUN dpkg -i cloudflared.deb
RUN echo "doppler run --command \"cloudflared tunnel run & lilypad solver --network ${network}\"" >> run

FROM base AS expose-local
EXPOSE 8080
RUN if [ "${network}" = "dev" ]; then \
        echo "lilypad solver --network $network" >> run; \
    else \
        echo "doppler run -- lilypad solver --network $network" >> run; \
    fi

FROM expose-$expose_via AS final
RUN chmod +x run

CMD ["/bin/bash", "./run"]
