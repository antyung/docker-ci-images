# syntax=docker/dockerfile:1

FROM docker.io/antyung/alpine-ci:3.20 AS base

ENV PNPM_HOME="/opt/pnpm" \
    POETRY_NO_INTERACTION=1 \
    POETRY_VIRTUALENVS_IN_PROJECT=1 \
    POETRY_VIRTUALENVS_CREATE=1 \
    POETRY_CACHE_DIR="/opt/poetry/.cache" \
    POETRY_HOME="/opt/poetry"

ENV PATH="${PNPM_HOME}:${POETRY_HOME}/bin:${POETRY_HOME}/venv/bin:${PATH}"

FROM docker.io/antyung/node:22.12.0-alpine AS node-build

ARG AWSCDK_VERSION=2.171.1

RUN pnpm install -g aws-cdk@${AWSCDK_VERSION}

FROM docker.io/antyung/python:3.13.0-alpine AS python-build

RUN source "${POETRY_HOME}/venv/bin/activate" \
    && poetry init --no-interaction --name aws-sam-cli-ci \
    && poetry add --no-interaction --no-cache \
    boto3 aws-sam-cli@latest

FROM base AS final

COPY --from=node-build "${PNPM_HOME}" "${PNPM_HOME}"
COPY --from=python-build "${POETRY_HOME}" "${POETRY_HOME}"

RUN apk add --no-cache nodejs npm \
    && rm -rf /var/cache/apk/*

USER root
