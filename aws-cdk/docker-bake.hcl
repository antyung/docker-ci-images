variable "DOCKERHUB_USERNAME" {
  default = "antyung"
}

variable "IMAGE" {
  default = "aws-cdk"
}

variable "TAG" {
  default = "latest"
}

group "default" {
  targets = ["build"]
}

target "settings" {
  context = "."
  cache-from = [
    "type=gha"
  ]
  cache-to = [
    "type=gha,mode=max"
  ]
}

target "test" {
  inherits = ["settings"]
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = []
}

target "build" {
  inherits = ["settings"]
  output   = ["type=docker"]
  tags = [
    "${DOCKERHUB_USERNAME}/${IMAGE}:latest",
    "${DOCKERHUB_USERNAME}/${IMAGE}:${TAG}",
  ]
}

target "push" {
  inherits = ["settings"]
  output   = ["type=registry"]
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  tags = [
    "${DOCKERHUB_USERNAME}/${IMAGE}:latest",
    "${DOCKERHUB_USERNAME}/${IMAGE}:${TAG}",
    "${DOCKERHUB_USERNAME}/${IMAGE}:${TAG}-alpine",
    "public.ecr.aws/w2u0w5i6/ci/${IMAGE}:latest",
    "public.ecr.aws/w2u0w5i6/ci/${IMAGE}:${TAG}",
    "public.ecr.aws/w2u0w5i6/ci/${IMAGE}:${TAG}-alpine",
  ]
}
