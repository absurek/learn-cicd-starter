
FROM ubuntu:24.04

ENV GOLANG_VERSION=1.25.5

# System setup
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl \
    netcat-openbsd \
    git \
    jq \
    wget \
    ca-certificates \
    xz-utils \
    postgresql-client && \
    rm -rf /var/lib/apt/lists/*

# Google Cloud SDK setup
RUN curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-linux-x86_64.tar.gz && \
    tar -xf google-cloud-cli-linux-x86_64.tar.gz && \
    ./google-cloud-sdk/install.sh -q && \
    rm google-cloud-cli-linux-x86_64.tar.gz

# Turso setup
RUN curl -sSfL https://get.tur.so/install.sh | bash

# Go setup
RUN wget -q https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

# GO CLIs
RUN go install github.com/bootdotdev/bootdev@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/securego/gosec/v2/cmd/gosec@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest && \
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
