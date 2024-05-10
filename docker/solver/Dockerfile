ARG expose_via=local

FROM golang:latest as base
WORKDIR /usr/src/app

ARG arch=amd64
ARG doppler_config=dev
ARG cloudflare_token="not-a-token"

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
RUN echo "cloudflared tunnel --metrics 0.0.0.0:11113 run --token $cloudflare_token --url http://localhost:8080 &" >> run

FROM base AS expose-local
EXPOSE 8080

FROM expose-$expose_via AS FINAL
RUN echo "doppler run -p solver -c $doppler_config -- lilypad solver" >> run
RUN chmod +x run

CMD ["/bin/bash", "./run"]
