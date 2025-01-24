package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Packer = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "packer",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "ci",
}

func TestBuildPacker(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", Packer.DOCKER_IMAGE),
			Dockerfile:    "Dockerfile",
			KeepImage:     false,
			PrintBuildLog: true,
		},
	}
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, e)
}

func TestPullPacker(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: fmt.Sprintf("%s/%s/%s:%s", Packer.AWS_ECR_URI, Packer.DOCKER_IMAGE_GROUP, Packer.DOCKER_IMAGE, Packer.DOCKER_TAG),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)
}
