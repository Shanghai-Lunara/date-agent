package date_agent

import (
	"context"
	"k8s.io/klog/v2"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const CommandTimeoutEnv = "DATE_AGENT_CMD_TIMEOUT"
const CommandTimeout = 15

func Exec(suffix []string) (out string, err error) {
	t := CommandTimeout
	if x := os.Getenv(CommandTimeoutEnv); x != "" {
		if t, err = strconv.Atoi(x); err != nil {
			klog.V(2).Info(err)
		} else {
			t = CommandTimeout
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(t))
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
