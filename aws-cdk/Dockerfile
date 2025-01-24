# syntax=docker/dockerfile:1

FROM public.ecr.aws/w2u0w5i6/base/python:3.13.1-alpine AS base

USER root

ENV PNPM_HOME="/opt/pnpm" \
    POETRY_NO_INTERACTION=1 \
    POETRY_VIRTUALENVS_IN_PROJECT=1 \
    POETRY_VIRTUALENVS_CREATE=1 \
    POETRY_CACHE_DIR="${HOME}/.cache/pypoetry" \
    POETRY_HOME="/opt/poetry"

ENV PATH="${PNPM_HOME}:${POETRY_HOME}/bin:${POETRY_HOME}/venv/bin:${PATH}"

RUN apk add --no-cache \
    sudo bash git aws-cli curl ca-certificates jq make nodejs npm \
    && rm -rf /var/cache/apk/* /etc/apk/cache/* \
    && ln -s /var/cache/apk /etc/apk/cache

FROM public.ecr.aws/w2u0w5i6/base/node:23.6.0-alpine AS build-node

ARG AWSCDK_VERSION=2.171.1

RUN pnpm install -g aws-cdk@${AWSCDK_VERSION}

FROM base AS build-python

WORKDIR "${POETRY_HOME}"

RUN curl -fsSL https://install.python-poetry.org | python3 -

COPY pyproject.toml .

RUN source "${POETRY_HOME}/venv/bin/activate" \
    && poetry install

FROM base AS final

COPY --from=build-node "${PNPM_HOME}" "${PNPM_HOME}"
COPY --from=build-python "${POETRY_HOME}" "${POETRY_HOME}"
