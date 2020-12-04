package main

import (
	"flag"
	agent "github.com/Shanghai-Lunara/date-agent"
	"github.com/nevercase/k8s-controller-custom-resource/pkg/signals"
	"k8s.io/klog/v2"
)

var (
	grpcservice string
)

func init() {
	flag.StringVar(&grpcservice, "grpcservice", "127.0.0.1:10000", "The address of the grpc server.")
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()
	client := agent.NewClient(grpcservice)
	<-stopCh
	client.Close()
}
