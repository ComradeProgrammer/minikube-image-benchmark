package command

import (
	"fmt"
	"os/exec"
)

// DockerSystemPrune does a docker system prune
func DockerSystemPrune() error {
	c := exec.Command("docker", "system", "prune", "-a", "--volumes", "-f")
	if _, err := run(c); err != nil {
		return fmt.Errorf("failed to docker prune: %v", err)
	}
	return nil
}

// minikubeDockerSystemPrune doese a minikube docker system prune
func minikubeDockerSystemPrune(profile string) error {
	args := fmt.Sprintf("eval $(./minikube -p %s docker-env) && docker system prune -a --volumes -f", profile)
	c := exec.Command("/bin/bash", "-c", args)
	if _, err := run(c); err != nil {
		return fmt.Errorf("failed to minikube docker prune: %v", err)
	}
	return nil
}
