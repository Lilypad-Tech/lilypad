FROM nvidia/cuda:12.6.3-devel-ubuntu22.04

ENV DEBIAN_FRONTEND=noninteractive \
    TZ=UTC \
    RUSTUP_HOME=/usr/local/rustup \
    CARGO_HOME=/usr/local/cargo \
    PATH=/usr/local/cargo/bin:$PATH \
    CUDA_ROOT=/usr/local/cuda \
    LD_LIBRARY_PATH=${CUDA_ROOT}/lib64:${LD_LIBRARY_PATH} \
    CUDA_COMPUTE_CAP=75 \
    NVIDIA_VISIBLE_DEVICES=all \
    NVIDIA_DRIVER_CAPABILITIES=gpu,compute,utility

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential pkg-config libssl-dev curl ca-certificates \
    nvidia-utils-560 cuda-cupti-12-6 libsasl2-dev cmake protobuf-compiler \
    libnvidia-ml-dev nvidia-container-toolkit git netcat tini \
    && rm -rf /var/lib/apt/lists/* \
    && ln -s /usr/lib/x86_64-linux-gnu/libnvidia-ml.so.1 /usr/lib/x86_64-linux-gnu/libnvidia-ml.so \
    && curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y \
    && . "/usr/local/cargo/env" \
    && rustup default stable

WORKDIR /workspace
COPY ../../pkg/gpu-tools/ .

# Build release version
RUN cargo build --release

# Use tini as entrypoint
ENTRYPOINT ["/usr/bin/tini", "--"]

# Run the binary through tini
CMD ["/workspace/target/release/gpu-tools"] 