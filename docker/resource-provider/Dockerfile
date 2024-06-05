FROM golang:latest as base
WORKDIR /usr/src/app

ARG doppler_config=dev
ARG DOPPLER_TOKEN_BACALHAU
ARG DOPPLER_TOKEN_RESOURCE_PROVIDER

COPY . .
RUN go mod download && go mod verify
RUN go build -v .
RUN go install

RUN (curl -Ls --tlsv1.2 --proto "=https" --retry 3 https://cli.doppler.com/install.sh || wget -t 3 -qO- https://cli.doppler.com/install.sh) | sh

# Install cron, used for PoW scheduler
RUN apt-get update && apt-get install -y cron

# TODO: introduce some form of randomisation here

# Add the cron job to run your script every n minutes
RUN echo "*/30 * * * * /usr/src/app/lilypad pow >> /var/log/cron.log 2>&1" > /etc/cron.d/pow-scheduler

# Give execution rights on the cron job
RUN chmod 0644 /etc/cron.d/pow-scheduler

# Apply the cron job
RUN crontab /etc/cron.d/pow-scheduler

# RUN wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.0.3/bacalhau_v1.0.3_linux_amd64.tar.gz
# RUN tar xfv bacalhau_v1.0.3_linux_amd64.tar.gz
# RUN mv bacalhau /usr/local/bin
# # RUN sysctl -w net.core.rmem_max=7500000
# # RUN sysctl -w net.core.wmem_max=7500000

# Create the run resource-provider script
RUN touch run
RUN echo "#!/bin/bash" >> run
#RUN echo "doppler run --token=$DOPPLER_TOKEN_BACALHAU -p bacalhau -c $doppler_config -- ./stack bacalhau-serve &" >> run
RUN echo "doppler run --token=$DOPPLER_TOKEN_RESOURCE_PROVIDER -p resource-provider -c $doppler_config -- lilypad resource-provider" >> run
RUN chmod +x run

# Create the run cron script
RUN echo "#!/bin/bash" > /usr/src/app/entrypoint.sh
RUN echo "cron" >> /usr/src/app/entrypoint.sh
RUN echo "/bin/bash /usr/src/app/run" >> /usr/src/app/entrypoint.sh
RUN chmod +x /usr/src/app/entrypoint.sh


ENTRYPOINT ["/usr/src/app/entrypoint.sh"]