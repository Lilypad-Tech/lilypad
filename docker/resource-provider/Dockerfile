FROM golang:latest as base
WORKDIR /usr/src/app

ARG doppler_config=dev
ARG network=dev
ARG DOPPLER_TOKEN_BACALHAU
ARG DOPPLER_TOKEN_RESOURCE_PROVIDER

COPY . .
RUN go mod download && go mod verify
RUN go build -v .
RUN go install

RUN (curl -Ls --tlsv1.2 --proto "=https" --retry 3 https://cli.doppler.com/install.sh || wget -t 3 -qO- https://cli.doppler.com/install.sh) | sh

# RUN wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.0.3/bacalhau_v1.0.3_linux_amd64.tar.gz
# RUN tar xfv bacalhau_v1.0.3_linux_amd64.tar.gz
# RUN mv bacalhau /usr/local/bin
# # RUN sysctl -w net.core.rmem_max=7500000
# # RUN sysctl -w net.core.wmem_max=7500000

RUN touch run
RUN echo "#!/bin/bash" >> run
#RUN echo "doppler run --token=$DOPPLER_TOKEN_BACALHAU -p bacalhau -c $doppler_config -- ./stack bacalhau-serve &" >> run
RUN echo "doppler run --token=$DOPPLER_TOKEN_RESOURCE_PROVIDER -p resource-provider -c $doppler_config -- lilypad resource-provider --network $network" >> run
RUN chmod +x run

CMD ["/bin/bash", "./run"]
