package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Ansible = struct {
	DOCKER_IMAGE       string
	DOCKER_TAG         string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "ansible",
	DOCKER_TAG:         "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "ci",
}

func TestBuildAnsible(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       fmt.Sprintf("../%s/", Ansible.DOCKER_IMAGE),
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

func TestPullAnsible(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: fmt.Sprintf("%s/%s/%s:%s", Ansible.AWS_ECR_URI, Ansible.DOCKER_IMAGE_GROUP, Ansible.DOCKER_IMAGE, Ansible.DOCKER_TAG),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err)
}
