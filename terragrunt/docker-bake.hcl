variable "AWS_ECR_PUBLIC_URI" {
  default = "public.ecr.aws/w2u0w5i6"
}

variable "GROUP" {
  default = "ci"
}

variable "IMAGE" {
  default = "terragrunt"
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
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:latest",
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${TAG}",
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
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:latest",
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${TAG}",
    "${AWS_ECR_PUBLIC_URI}/${GROUP}/${IMAGE}:${TAG}-alpine",
  ]
}
