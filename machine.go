package fns

import (
	"context"
	"log"
	"strings"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/yuekcc/fns/fnspec"
)

type Machine struct {
	cli *client.Client
	sync.Mutex
}

func GetClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}

func (r *Machine) useClient() *client.Client {
	r.Lock()
	defer r.Unlock()

	if r.cli == nil {
		cli, err := GetClient()
		if err != nil {
			panic(err)
		}

		r.cli = cli
	}

	return r.cli
}

func (r *Machine) createContainer(ctx context.Context, imageName string, labels map[string]string) (string, error) {
	config := &container.Config{
		Image:  imageName,
		Labels: labels,
	}

	resp, err := r.useClient().ContainerCreate(ctx, config, nil, nil, nil, "")
	if err != nil {
		log.Println("ContainerCreate Error:", err)
		return "", err
	}

	return resp.ID, nil
}

func (r *Machine) Install(ctx context.Context, spec *fnspec.FunctionSpec) (string, error) {
	entryPoint := spec.Metadata.LaunchCommandLine
	if entryPoint == "" {
		entryPoint = spec.Metadata.Name
	}

	labels := generateContainerLabels(spec)
	imageName := "yuekcc/node-demo-app"
	return r.createContainer(ctx, imageName, labels)
}

func (r *Machine) Spawn(ctx context.Context, containerID string) error {
	err := r.useClient().ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		log.Println("ContainerStart Error:", err)
		return err
	}

	return nil
}

func (r *Machine) Kill(ctx context.Context, containerID string) error {
	err := r.useClient().ContainerStop(ctx, containerID, nil)
	if err != nil {
		log.Println("ContainerStop Error:", err)
		return err
	}

	return nil
}

func (r *Machine) Uninstall(ctx context.Context, containerID string) error {
	info, err := r.useClient().ContainerInspect(ctx, containerID)
	if err != nil {
		log.Println("ContainerStats Error:", err)
		return err
	}

	containerStatus := strings.ToLower(info.State.Status)
	if containerStatus == "running" {
		log.Println("killing a running app")
		err := r.Kill(ctx, containerID)
		if err != nil {
			return err
		}
	}

	err = r.useClient().ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err != nil {
		log.Println("ContainerRemove Error:", err)
		return err
	}

	return nil
}

func (r *Machine) List(ctx context.Context) ([]types.Container, error) {
	list, err := r.useClient().ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	result := []types.Container{}
	for _, c := range list {
		_, exists := c.Labels["fns-spec-version"] // 只显示包含 fns-spec-version 标签的容器
		if exists {
			result = append(result, c)
		}
	}

	return result, nil
}

func (m *Machine) Inspect(ctx context.Context, containerId string) (types.ContainerJSON, error) {
	return m.useClient().ContainerInspect(ctx, containerId)
}

func (r *Machine) Shutdown() error {
	return r.useClient().Close()
}

func GetMachine() *Machine {
	return &Machine{}
}
