package repository

type DockerRepository interface {
	RestartContainer(containerName string) error
	Connect(endpoint string) error
}