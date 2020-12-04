package main

import (
	"flag"
	agent "github.com/Shanghai-Lunara/date-agent"
	"github.com/nevercase/k8s-controller-custom-resource/pkg/signals"
	"k8s.io/klog/v2"
)

var (
	grpcservice string
	httpservice string
)

func init() {
	flag.StringVar(&grpcservice, "grpcservice", "0.0.0.0:10000", "The address of the grpc server.")
	flag.StringVar(&httpservice, "httpservice", "0.0.0.0:10001", "The address of the http server.")
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()
	server := agent.NewServer(grpcservice, httpservice)
	<-stopCh
	server.Close()
}
