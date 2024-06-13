FROM ubuntu:latest
WORKDIR /usr/src/app

# Default environment variables
ENV LOG_TYPE=json
ENV LOG_LEVEL=debug
ENV HOME=/app/lilypad
ENV BACALHAU_SERVE_IPFS_PATH=/data/ipfs
ENV OFFER_GPU=1

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    git \
    curl \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Installing Bacalhau
RUN cd /tmp && \
    wget https://github.com/bacalhau-project/bacalhau/releases/download/v1.3.2/bacalhau_v1.3.2_linux_amd64.tar.gz && \
    tar xfv bacalhau_v1.3.2_linux_amd64.tar.gz && \
    mv bacalhau /usr/bin/bacalhau && \
    rm bacalhau_v1.3.2_linux_amd64.tar.gz
# COPY bacalhau /usr/bin/bacalhau
# RUN  mkdir -p /app/data/ipfs

# Installing Lilypad
RUN OSARCH=$(uname -m | awk '{if ($0 ~ /arm64|aarch64/) print "arm64"; else if ($0 ~ /x86_64|amd64/) print "amd64"; else print "unsupported_arch"}') && \
    OSNAME=$(uname -s | awk '{if ($1 == "Darwin") print "darwin"; else if ($1 == "Linux") print "linux"; else print "unsupported_os"}') && \
    curl https://api.github.com/repos/lilypad-tech/lilypad/releases/latest | \
    grep "browser_download_url.*lilypad-$OSNAME-$OSARCH" | \
    cut -d : -f 2,3 | tr -d \" | wget -qi - -O lilypad && \
    chmod +x lilypad && \
    mv lilypad /usr/local/bin/lilypad

# COPY lilypad lilypad
# RUN chmod +x lilypad && \
#     mv lilypad /usr/local/bin/lilypad

# Add both lilypad and bacalhau executables to PATH
ENV PATH="/usr/local/bin:/usr/bin:${PATH}"

# Create a startup script to run both services simultaneously
RUN touch run

RUN echo "#!/bin/bash" >> run
# Launch Bacalhau
RUN echo "/usr/bin/bacalhau serve --node-type compute,requester --peer none --private-internal-ipfs=false &" >> run
# Launch Lilypad
RUN echo "/usr/local/bin/lilypad resource-provider &" >> run
RUN echo "wait -n" >> run
RUN chmod +x run

# Run startup script when container starts
CMD ["/bin/bash", "./run"]