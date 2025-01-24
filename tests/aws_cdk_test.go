package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var AwsCdk = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "aws-cdk",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "ci",
}

func TestBuildAwsCdk(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", AwsCdk.DOCKER_IMAGE),
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

func TestPullAwsCdk(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: fmt.Sprintf("%s/%s/%s:%s", AwsCdk.AWS_ECR_URI, AwsCdk.DOCKER_IMAGE_GROUP, AwsCdk.DOCKER_IMAGE, AwsCdk.DOCKER_TAG),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)
}
