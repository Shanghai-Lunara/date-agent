package date_agent

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"k8s.io/klog"
)

// getHostName : get the hostname of the host machine if the container is started by docker run --net=host
func getHostName() (string, error) {
	cmd := exec.Command("/bin/hostname")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	hostname := strings.TrimSpace(string(out))
	if hostname == "" {
		return "", fmt.Errorf("no hostname get from cmd '/bin/hostname' in the container, please check")
	}
	return hostname, nil
}

// GetHostName : get the hostname of host machine
func GetHostName() (string, error) {
	hostName := os.Getenv("HOST_NAME")
	if hostName != "" {
		return hostName, nil
	}
	klog.Info("get HOST_NAME from env failed, is env.(\"HOST_NAME\") already set? Will use hostname instead")
	return getHostName()
}
