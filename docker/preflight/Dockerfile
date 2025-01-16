# FROM alpine:latest AS base
# WORKDIR /usr/src/app

# ARG network=dev

# COPY . .
# RUN go mod download && go mod verify
# RUN go build -v .
# RUN go install

# RUN (curl -Ls --tlsv1.2 --proto "=https" --retry 3 https://cli.doppler.com/install.sh || wget -t 3 -qO- https://cli.doppler.com/install.sh) | sh

# RUN touch run
# RUN echo "#!/bin/bash" >> run
# RUN if [ "${network}" = "dev" ]; then \
#         echo "lilypad jobcreator --network $network" >> run; \
#     else \
#         echo "doppler run -- lilypad jobcreator --network $network" >> run; \
#     fi
# RUN chmod +x run
# WORKDIR /app
# COPY ./preflight /app
# RUN chmod +x /app/preflight
# # CMD ["/app/preflight"]
# # CMD ["/bin/sh"]
# CMD  tail -f /dev/null
FROM golang:1.23 AS builder
WORKDIR /app
# RUN pwd 
COPY ./preflight.go /app
# COPY . .
# RUN cd preflight
RUN go mod init preflight
RUN go mod tidy
RUN go build -o ./preflight preflight.go

FROM nvidia/cuda:12.0.1-cudnn8-devel-ubuntu22.04 AS base

RUN apt-get update && \
    apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release && \
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - && \
    echo "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null && \
    apt-get update && \
    apt-get install -y docker-ce docker-ce-cli containerd.io && \
    rm -rf /var/lib/apt/lists/*


WORKDIR /app
COPY --from=builder /app .
RUN chmod +x /app/preflight
# CMD ["/app/preflight","arsenum"]