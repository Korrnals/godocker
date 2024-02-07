package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Структура Containers содержит информацию о контейнере.
type Containers struct {
	Name    []string
	ID      string
	Status  string
	Created string
	Ports   []string
	IP      string
	Mounts  []types.MountPoint
	Labels  map[string]string
	Image   string
}

func main() {
	c := Containers{}
	c.ListAll()
}

// Метод ListAll структуры Containers выводит список всех контейнеров и их параметры в структуре Containers.
func (c *Containers) ListAll() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, containerInfo := range containers {
		container := c.mapContainerInfoToStruct(containerInfo)
		c.printContainerInfo(container)
	}
}

// Метод mapContainerInfoToStruct структуры Containers для преобразования информации о контейнере в структуру Containers.
func (c *Containers) mapContainerInfoToStruct(info types.Container) Containers {
	container := Containers{
		Name:    info.Names,
		ID:      info.ID,
		Status:  info.Status,
		Created: c.formatTimeSinceCreation(info),
		Mounts:  info.Mounts,
		Labels:  info.Labels,
		Image:   info.Image,
	}

	for _, settings := range info.NetworkSettings.Networks { // поиск IP адреса контейнера
		container.IP = settings.IPAddress

		var ports []string // слайс для форматирования портов контейнера

		for _, port := range info.Ports {
			portStr := fmt.Sprintf(
				"{Container port: %v, Host port: %v}",
				port.PrivatePort, port.PublicPort) // форматирование портов

			ports = append(ports, portStr) // добавление портов в список
		}
		container.Ports = ports // присвоение списка портов контейнеру в структуре Containers
	}

	return container
}

// Метод formatTimeSinceCreation структуры Containers для форматирования времени создания контейнера.
func (c *Containers) formatTimeSinceCreation(container types.Container) string {
	createTime := time.Unix(container.Created, 0)
	timeSinceCreation := time.Now().Sub(createTime)

	days := int(timeSinceCreation.Hours() / 24)
	if days > 0 {
		return fmt.Sprintf("%d days ago", days)
	}
	return fmt.Sprintf("Today")
}

// Метод printContainerInfo структуры Containers для вывода информации о контейнере.
func (c *Containers) printContainerInfo(container Containers) {
	for _, name := range container.Name { // цикл для очищения имени контейнера от скобок
		cleanedName := strings.Trim(name, "[/]")
		fmt.Printf("Name: %s\n", cleanedName)
		fmt.Printf("ID: %s\n", container.ID[:12])
		fmt.Printf("Image: %s\n", container.Image)
		fmt.Printf("Status: %s\n", container.Status)
		fmt.Printf("Created: %s\n", container.Created)
		if container.IP != "" { // условие для вывода информации, если контейнер имеет сетевые настройки
			fmt.Printf("Container Net: %s/%d\n", container.IP, 24)
			fmt.Printf("Container IP: %v\n", container.IP)
			fmt.Printf("Ports: %v\n", container.Ports)
		}
		fmt.Println()
	}
}
