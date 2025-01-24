package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Kubectl = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "kubectl",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "ci",
}

func TestBuildKubectl(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + Kubectl.DOCKER_IMAGE + "/",
				Dockerfile:    "Dockerfile",
				KeepImage:     false,
				PrintBuildLog: true,
			},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestPullKubectl(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Kubectl.AWS_ECR_URI + "/" + Kubectl.DOCKER_IMAGE_GROUP + "/" + Kubectl.DOCKER_IMAGE + ":" + Kubectl.DOCKER_TAG,
		},
		Started: false,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}
