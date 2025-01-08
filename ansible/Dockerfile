# syntax=docker/dockerfile:1

FROM public.ecr.aws/w2u0w5i6/base/python:3.13.1-alpine AS base

ENV POETRY_NO_INTERACTION=1 \
    POETRY_VIRTUALENVS_IN_PROJECT=1 \
    POETRY_VIRTUALENVS_CREATE=1 \
    POETRY_CACHE_DIR="/opt/poetry/.cache" \
    POETRY_HOME="/opt/poetry"

ENV PATH="${POETRY_HOME}/bin:${POETRY_HOME}/venv/bin:${PATH}"

RUN apk add --no-cache \
    sudo bash git aws-cli curl ca-certificates jq make openssh sshpass libc6-compat \
    && rm -rf /var/cache/apk/* /etc/apk/cache/* \
    && ln -s /var/cache/apk /etc/apk/cache

FROM base AS build

WORKDIR "${POETRY_HOME}"

RUN case `uname -m` in \
    x86_64) ARCH=64bit; ;; \
    aarch64) ARCH=arm64; ;; \
    *) echo "un-supported arch, exit ..."; exit 1; ;; \
    esac \
    && echo "${ARCH}" > /tmp/ARCH

RUN apk add --no-cache dpkg \
    && curl -fsSL https://install.python-poetry.org | python3 -

COPY pyproject.toml .

RUN source "${POETRY_HOME}/venv/bin/activate" \
    && poetry install

RUN export ARCH=$(cat /tmp/ARCH) \
    && curl -fsSL "https://s3.amazonaws.com/session-manager-downloads/plugin/latest/ubuntu_${ARCH}/session-manager-plugin.deb" -o /tmp/session-manager-plugin.deb \
    && dpkg -x /tmp/session-manager-plugin.deb /tmp/session-manager-plugin \
    && mv /tmp/session-manager-plugin/usr/local/sessionmanagerplugin/bin/session-manager-plugin /usr/local/bin/session-manager-plugin

FROM base AS final

COPY --from=build "${POETRY_HOME}" "${POETRY_HOME}"
COPY --from=build /usr/local/bin/session-manager-plugin /usr/local/bin/session-manager-plugin

RUN chmod +x /usr/local/bin/session-manager-plugin

USER root
