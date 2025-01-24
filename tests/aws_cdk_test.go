package tests

import (
	"context"
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
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + AwsCdk.DOCKER_IMAGE + "/",
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

func TestPullAwsCdk(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: AwsCdk.AWS_ECR_URI + "/" + AwsCdk.DOCKER_IMAGE_GROUP + "/" + AwsCdk.DOCKER_IMAGE + ":" + AwsCdk.DOCKER_TAG,
		},
		Started: false,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}
