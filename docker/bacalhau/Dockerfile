FROM nvidia/cuda:12.0.1-cudnn8-runtime-ubuntu22.04 

WORKDIR /app

RUN apt update && \
  apt install -y apt-utils bash ca-certificates curl gnupg iptables && \
  apt clean

# Install NVIDIA CTK and Docker 
RUN mkdir -pm755 /etc/apt/keyrings && curl -o /etc/apt/keyrings/docker.asc -fsSL "https://download.docker.com/linux/ubuntu/gpg" && chmod a+r /etc/apt/keyrings/docker.asc && \
  mkdir -pm755 /etc/apt/sources.list.d && echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(grep UBUNTU_CODENAME= /etc/os-release | cut -d= -f2 | tr -d '\"') stable" > /etc/apt/sources.list.d/docker.list && \
  mkdir -pm755 /usr/share/keyrings && curl -fsSL "https://nvidia.github.io/libnvidia-container/gpgkey" | gpg --dearmor -o /usr/share/keyrings/nvidia-container-toolkit-keyring.gpg && \
  curl -fsSL "https://nvidia.github.io/libnvidia-container/stable/deb/nvidia-container-toolkit.list" | sed 's#deb https://#deb [signed-by=/usr/share/keyrings/nvidia-container-toolkit-keyring.gpg] https://#g' > /etc/apt/sources.list.d/nvidia-container-toolkit.list && \
  apt-get update && \
  apt-get install -y nvidia-container-toolkit docker-ce docker-ce-cli && \
  apt clean
RUN nvidia-ctk runtime configure --runtime=docker --set-as-default

# Install Bacalhau
ADD https://github.com/bacalhau-project/bacalhau/releases/download/v1.6.0/bacalhau_v1.6.0_linux_amd64.tar.gz .
RUN tar xfv bacalhau_v1.6.0_linux_amd64.tar.gz
RUN mv bacalhau /usr/local/bin

ADD https://raw.githubusercontent.com/moby/moby/refs/heads/master/hack/dind /usr/local/bin/dind
RUN chmod +x /usr/local/bin/dind

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=10 \
  CMD curl http://localhost:1234/api/v1/agent/alive -q || exit 1

VOLUME /var/lib/docker

RUN touch run
RUN echo "#!/bin/bash" >> run
RUN echo "find /run /var/run -iname 'docker*.pid' -delete" >> run
RUN echo "dind dockerd &" >> run
RUN echo 'until pgrep "dockerd" >/dev/null; do sleep 0.5; done' >> run
RUN echo 'exec "$@"' >> run
RUN chmod a+x ./run

ENTRYPOINT [ "./run" ]