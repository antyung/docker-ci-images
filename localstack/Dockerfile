# syntax=docker/dockerfile:1

FROM localstack/localstack:4.0.3 AS base

USER root

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get install -y --no-install-recommends locales-all \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

ENV TZ="Asia/Singapore" \
    LANG="en_US.UTF-8" \
    LANGUAGE="en_US:en" \
    LC_ALL="en_US.UTF-8"
