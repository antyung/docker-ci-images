variable "AWS_ECR_PUBLIC_URI" {
  default = "public.ecr.aws/w2u0w5i6"
}

variable "GROUP" {
  default = "ci"
}

variable "IMAGE" {
  default = "ansible"
}

variable "DOCKER_TAG" {
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
  DOCKER_TAGs = []
}

target "build" {
  inherits = ["settings"]
  output   = ["type=docker"]
  DOCKER_TAGs = [
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:latest",
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${DOCKER_TAG}",
  ]
}

target "push" {
  inherits = ["settings"]
  output   = ["type=registry"]
  platforms = [
    "linux/amd64",
    "linux/arm64",
  ]
  DOCKER_TAGs = [
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:latest",
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${DOCKER_TAG}",
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${DOCKER_TAG}-alpine",
  ]
}
