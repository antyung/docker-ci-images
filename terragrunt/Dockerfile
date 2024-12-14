# syntax=docker/dockerfile:1

FROM public.ecr.aws/w2u0w5i6/base/alpine-ci:3.21 AS base

FROM base AS build

ARG TERRAGRUNT_VERSION=0.69.1
ARG TERRAFORM_VERSION=1.10.0
ARG TERRAFORMDOCS_VERSION=0.19.0

RUN case `uname -m` in \
    x86_64) ARCH=linux_amd64; ;; \
    aarch64) ARCH=linux_arm64; ;; \
    *) echo "un-supported arch, exit ..."; exit 1; ;; \
    esac \
    && echo "${ARCH}" > /tmp/ARCH

RUN apk add unzip tar

RUN export ARCH=$(cat /tmp/ARCH) \
    && curl -fsSL "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${ARCH}.zip" -o "/tmp/terraform_${TERRAFORM_VERSION}_${ARCH}.zip" \
    && unzip "/tmp/terraform_${TERRAFORM_VERSION}_${ARCH}.zip" -d /tmp \
    && mv /tmp/terraform /usr/local/bin/terraform

RUN export ARCH=$(cat /tmp/ARCH) \
    && curl -fsSL "https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_${ARCH}" -o "/tmp/terragrunt_${ARCH}" \
    && mv "/tmp/terragrunt_${ARCH}" /usr/local/bin/terragrunt

RUN export ARCH=$(cat /tmp/ARCH | sed 's/_/-/g') \
    && curl -fsSL "https://github.com/terraform-docs/terraform-docs/releases/download/v${TERRAFORMDOCS_VERSION}/terraform-docs-v${TERRAFORMDOCS_VERSION}-${ARCH}.tar.gz" -o "/tmp/terraform-docs-v${TERRAFORMDOCS_VERSION}-${ARCH}.tar.gz" \
    && tar -xvf "/tmp/terraform-docs-v${TERRAFORMDOCS_VERSION}-${ARCH}.tar.gz" -C /tmp \
    && mv /tmp/terraform-docs /usr/local/bin/terraform-docs

FROM base AS final

COPY --from=build /usr/local/bin/terraform /usr/local/bin/terraform
COPY --from=build /usr/local/bin/terragrunt /usr/local/bin/terragrunt
COPY --from=build /usr/local/bin/terraform-docs /usr/local/bin/terraform-docs

RUN chmod +x /usr/local/bin/terraform /usr/local/bin/terragrunt /usr/local/bin/terraform-docs

USER root
