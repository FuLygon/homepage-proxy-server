package docker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Docker interface {
	Exec(ctx context.Context, targetContainer string, command []string) (*bytes.Buffer, error)
}
type docker struct {
	client *client.Client
}

func NewDocker(client *client.Client) Docker {
	return &docker{
		client: client,
	}
}

func (d *docker) Exec(ctx context.Context, targetContainer string, command []string) (*bytes.Buffer, error) {
	// Create an exec instance
	execConfig := container.ExecOptions{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          command,
	}
	exec, err := d.client.ContainerExecCreate(ctx, targetContainer, execConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating exec instance: %w", err)
	}
	execID := exec.ID

	// Start exec instance
	attachResp, err := d.client.ContainerExecAttach(ctx, execID, container.ExecAttachOptions{})
	if err != nil {
		return nil, fmt.Errorf("error attaching to exec instance: %w", err)
	}
	defer attachResp.Close()

	// Read the output
	var stdout, stderr bytes.Buffer
	_, err = stdcopy.StdCopy(&stdout, &stderr, attachResp.Reader)
	if err != nil {
		return nil, fmt.Errorf("error copying exec output: %w", err)
	}

	if len(stderr.Bytes()) > 0 {
		return nil, errors.New(stderr.String())
	}

	// Check the exit code
	inspectResp, err := d.client.ContainerExecInspect(ctx, execID)
	if err != nil {
		return nil, fmt.Errorf("error inspecting exec instance: %w", err)
	}

	if inspectResp.ExitCode != 0 {
		return nil, fmt.Errorf("command exited with code %d", inspectResp.ExitCode)
	}

	return &stdout, nil
}
