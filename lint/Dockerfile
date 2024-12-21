# syntax=docker/dockerfile:1

FROM public.ecr.aws/w2u0w5i6/base/python:3.13.1-alpine AS base

ENV POETRY_NO_INTERACTION=1 \
    POETRY_VIRTUALENVS_IN_PROJECT=1 \
    POETRY_VIRTUALENVS_CREATE=1 \
    POETRY_CACHE_DIR="/opt/poetry/.cache" \
    POETRY_HOME="/opt/poetry"

ENV PATH="${POETRY_HOME}/bin:${POETRY_HOME}/venv/bin:${PATH}"

RUN apk add --no-cache \
    sudo bash git aws-cli curl ca-certificates jq \
    && rm -rf /var/cache/apk/* /etc/apk/cache/* \
    && mkdir -p /var/cache/apk \
    && ln -s /var/cache/apk /etc/apk/cache

FROM base AS build

WORKDIR "${POETRY_HOME}"

RUN curl -fsSL https://install.python-poetry.org | python3 -

COPY pyproject.toml .

RUN source "${POETRY_HOME}/venv/bin/activate" \
    && poetry install

FROM base AS final

COPY --from=build "${POETRY_HOME}" "${POETRY_HOME}"

USER root
