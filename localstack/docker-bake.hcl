variable "DOCKER_IMAGE" {
  default = "localstack"
}

variable "DOCKER_TAG" {
  default = "latest"
}

variable "AWS_ECR_URI" {
  default = "public.ecr.aws/w2u0w5i6"
}

variable "DOCKER_IMAGE_GROUP" {
  default = "ci"
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
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:latest",
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_TAG}",
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
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:latest",
    "${AWS_ECR_URI}/${DOCKER_IMAGE_GROUP}/${DOCKER_IMAGE}:${DOCKER_TAG}",
  ]
}
