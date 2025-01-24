# syntax=docker/dockerfile:1

FROM localstack/localstack:4.0.3 AS base

USER root

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get install -y --no-install-recommends locales-all awscli \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

ENV TZ="Asia/Singapore" \
    LANG="en_US.UTF-8" \
    LANGUAGE="en_US:en" \
    LC_ALL="en_US.UTF-8"

FROM base AS build

ARG TERRAGRUNT_VERSION=0.69.1
ARG TERRAFORM_VERSION=1.10.0

RUN case `uname -m` in \
    x86_64) ARCH=linux_amd64; ;; \
    aarch64) ARCH=linux_arm64; ;; \
    *) echo "un-supported arch, exit ..."; exit 1; ;; \
    esac \
    && echo "${ARCH}" > /tmp/ARCH

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get install -y --no-install-recommends unzip tar

RUN export ARCH=$(cat /tmp/ARCH) \
    && curl -fsSL "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${ARCH}.zip" -o "/tmp/terraform_${TERRAFORM_VERSION}_${ARCH}.zip" \
    && unzip "/tmp/terraform_${TERRAFORM_VERSION}_${ARCH}.zip" -d /tmp \
    && cp /tmp/terraform /usr/local/bin/terraform

RUN export ARCH=$(cat /tmp/ARCH) \
    && curl -fsSL "https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_${ARCH}" -o "/tmp/terragrunt_${ARCH}" \
    && cp "/tmp/terragrunt_${ARCH}" /usr/local/bin/terragrunt

FROM base AS final

COPY --from=build /usr/local/bin/terraform /usr/local/bin/terraform
COPY --from=build /usr/local/bin/terragrunt /usr/local/bin/terragrunt

RUN chmod +x /usr/local/bin/terraform /usr/local/bin/terragrunt
