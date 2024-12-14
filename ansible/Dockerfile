# syntax=docker/dockerfile:1

FROM public.ecr.aws/w2u0w5i6/base/python:3.13.1-alpine AS base

RUN apk add --no-cache \
    sudo bash git aws-cli curl ca-certificates jq make openssh sshpass libc6-compat \
    && rm -rf /var/cache/apk/* /etc/apk/cache/* \
    && mkdir -p /var/cache/apk \
    && ln -s /var/cache/apk /etc/apk/cache

FROM base AS build

ARG ANSIBLE_VERSION=10.6.0

WORKDIR "${POETRY_HOME}"

RUN case `uname -m` in \
    x86_64) ARCH=64bit; ;; \
    aarch64) ARCH=arm64; ;; \
    *) echo "un-supported arch, exit ..."; exit 1; ;; \
    esac \
    && echo "${ARCH}" > /tmp/ARCH

RUN apk add dpkg \
    && source "${POETRY_HOME}/venv/bin/activate" \
    && poetry init --no-interaction --name ansible-ci \
    && poetry add --no-interaction --no-cache \
    boto3 ansible@${ANSIBLE_VERSION}

RUN export ARCH=$(cat /tmp/ARCH) \
    && curl -fsSL "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/ubuntu_${ARCH}/session-manager-plugin.deb" -o /tmp/session-manager-plugin.deb \
    && dpkg -x /tmp/session-manager-plugin.deb /tmp/session-manager-plugin \
    && mv /tmp/session-manager-plugin/usr/local/sessionmanagerplugin/bin/session-manager-plugin /usr/local/bin/session-manager-plugin

FROM base AS final

COPY --from=build "${POETRY_HOME}" "${POETRY_HOME}"
COPY --from=build /usr/local/bin/session-manager-plugin /usr/local/bin/session-manager-plugin

RUN chmod +x /usr/local/bin/session-manager-plugin

USER root
