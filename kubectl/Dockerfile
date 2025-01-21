# syntax=docker/dockerfile:1

FROM public.ecr.aws/w2u0w5i6/base/alpine-ci:3.21 AS base

USER root

FROM base AS build

ARG HELM_VERSION=3.17.0
ARG KUBECTL_VERSION=1.32.0

RUN case `uname -m` in \
    x86_64) ARCH=amd64; ;; \
    aarch64) ARCH=arm64; ;; \
    *) echo "un-supported arch, exit ..."; exit 1; ;; \
    esac \
    && echo "${ARCH}" > /tmp/ARCH

RUN apk add --no-cache tar

RUN export ARCH=$(cat /tmp/ARCH) \
    && curl -fsSL "https://get.helm.sh/helm-v${HELM_VERSION}-rc.1-linux-${ARCH}.tar.gz" -o "/tmp/helm-v${HELM_VERSION}-linux-${ARCH}.tar.gz" \
    && tar -xvf "/tmp/helm-v${HELM_VERSION}-linux-${ARCH}.tar.gz" -C /tmp \
    && cp "/tmp/linux-${ARCH}/helm" /usr/local/bin

RUN export ARCH=$(cat /tmp/ARCH) \
    && curl -fsSL "https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/${ARCH}/kubectl" -o "/usr/local/bin/kubectl"

FROM base AS final

COPY --from=build /usr/local/bin/helm /usr/local/bin/helm
COPY --from=build /usr/local/bin/kubectl /usr/local/bin/kubectl

RUN chmod +x /usr/local/bin/helm /usr/local/bin/kubectl

RUN helm completion bash > /etc/bash_completion.d/helm \
    && kubectl completion bash > /etc/bash_completion.d/kubectl
