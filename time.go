package date_agent

import (
	"context"
	"k8s.io/klog/v2"
	"os/exec"
	"time"
)

func Exec(suffix []string) (out string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	var res []byte
	commands := []string{"-c"}
	commands = append(commands, suffix...)
	res, err = exec.CommandContext(ctx, "/bin/bash", commands...).Output()
	cancel()
	if err != nil {
		klog.V(2).Info(err)
		return "", err
	}
	return string(res), nil
}
