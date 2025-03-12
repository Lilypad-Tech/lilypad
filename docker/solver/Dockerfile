FROM golang:latest AS base

FROM base AS build
WORKDIR /usr/src/app

COPY . .
RUN go mod download && go mod verify
RUN go build -v -ldflags="-X 'github.com/lilypad-tech/lilypad/pkg/system.Version=${VERSION}' -X 'github.com/lilypad-tech/lilypad/pkg/system.CommitSHA=${COMMIT_SHA}'"

FROM base AS final 
COPY --from=build /usr/src/app/lilypad /usr/local/bin

RUN (curl -Ls --tlsv1.2 --proto "=https" --retry 3 https://cli.doppler.com/install.sh || wget -t 3 -qO- https://cli.doppler.com/install.sh) | sh

RUN ARCH="$(dpkg --print-architecture)"; curl -L --output cloudflared.deb https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-${ARCH}.deb
RUN dpkg -i cloudflared.deb

EXPOSE 8080
COPY ./docker/solver/entrypoint.sh /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]