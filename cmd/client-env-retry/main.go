package main

import (
	"flag"
	agent "github.com/Shanghai-Lunara/date-agent"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	agent.NewClientByEnv()
}
