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

func getDockerClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}

func (m *Machine) useClient() *client.Client {
	m.Lock()
	defer m.Unlock()

	if m.cli == nil {
		cli, err := getDockerClient()
		if err != nil {
			panic(err)
		}

		m.cli = cli
	}

	return m.cli
}

func (m *Machine) createContainer(ctx context.Context, imageName string, labels map[string]string) (string, error) {
	config := &container.Config{
		Image:  imageName,
		Labels: labels,
	}

	resp, err := m.useClient().ContainerCreate(ctx, config, nil, nil, nil, "")
	if err != nil {
		log.Println("ContainerCreate Error:", err)
		return "", err
	}

	return resp.ID, nil
}

func (m *Machine) Install(ctx context.Context, spec *fnspec.FunctionSpec) (string, error) {
	labels := generateContainerLabels(spec)
	imageName := "yuekcc/node-demo-app"
	return m.createContainer(ctx, imageName, labels)
}

func (m *Machine) Spawn(ctx context.Context, containerId string) error {
	err := m.useClient().ContainerStart(ctx, containerId, types.ContainerStartOptions{})
	if err != nil {
		log.Println("ContainerStart Error:", err)
		return err
	}

	return nil
}

func (m *Machine) Kill(ctx context.Context, containerId string) error {
	err := m.useClient().ContainerStop(ctx, containerId, nil)
	if err != nil {
		log.Println("ContainerStop Error:", err)
		return err
	}

	return nil
}

func (m *Machine) Uninstall(ctx context.Context, containerId string) error {
	info, err := m.Inspect(ctx, containerId)
	if err != nil {
		log.Println("ContainerStats Error:", err)
		return err
	}

	containerStatus := strings.ToLower(info.State.Status)
	if containerStatus == "running" {
		log.Println("killing a running app")
		err := m.Kill(ctx, containerId)
		if err != nil {
			return err
		}
	}

	err = m.useClient().ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{})
	if err != nil {
		log.Println("ContainerRemove Error:", err)
		return err
	}

	return nil
}

func (m *Machine) List(ctx context.Context) ([]types.Container, error) {
	list, err := m.useClient().ContainerList(ctx, types.ContainerListOptions{All: true})
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

func (m *Machine) Shutdown() error {
	return m.useClient().Close()
}

func GetMachine() *Machine {
	return &Machine{}
}
