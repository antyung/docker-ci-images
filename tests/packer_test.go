package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Packer = struct {
	DOCKER_IMAGE       string
	DOCKER_IMAGE_TAG   string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "packer",
	DOCKER_IMAGE_TAG:   "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "ci",
}

func TestContainerBuildPacker(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + Packer.DOCKER_IMAGE + "/",
				Dockerfile:    "Dockerfile.alpine",
				KeepImage:     false,
				PrintBuildLog: true,
			},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestContainerPullPacker(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Packer.AWS_ECR_URI + "/" + Packer.DOCKER_IMAGE_GROUP + "/" + Packer.DOCKER_IMAGE + ":" + Packer.DOCKER_IMAGE_TAG,
		},
		Started: false,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestContainerExecPacker(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Packer.AWS_ECR_URI + "/" + Packer.DOCKER_IMAGE_GROUP + "/" + Packer.DOCKER_IMAGE + ":" + Packer.DOCKER_IMAGE_TAG,
			Cmd:   []string{"packer", "--version"},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}
