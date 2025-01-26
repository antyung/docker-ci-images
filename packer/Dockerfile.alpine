# syntax=docker/dockerfile:1

FROM public.ecr.aws/w2u0w5i6/base/alpine-ci:3.21 AS base

USER root

FROM base AS build

ARG PACKER_VERSION=1.11.2

RUN case `uname -m` in \
    x86_64) ARCH=amd64; ;; \
    aarch64) ARCH=arm64; ;; \
    *) echo "un-supported arch, exit ..."; exit 1; ;; \
    esac \
    && echo "${ARCH}" > /tmp/ARCH

RUN apk add --no-cache unzip

RUN export ARCH=$(cat /tmp/ARCH) \
    && curl -fsSL "https://releases.hashicorp.com/packer/${PACKER_VERSION}/packer_${PACKER_VERSION}_linux_${ARCH}.zip" -o "/tmp/packer_${PACKER_VERSION}_linux_${ARCH}.zip" \
    && unzip "/tmp/packer_${PACKER_VERSION}_linux_${ARCH}.zip" -d /tmp \
    && cp /tmp/packer /usr/local/bin/packer

FROM base AS final

COPY --from=build /usr/local/bin/packer /usr/local/bin/packer

RUN chmod +x /usr/local/bin/packer
