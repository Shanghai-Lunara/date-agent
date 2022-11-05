package main

import (
	"github.com/TyrandeCloud/signals/pkg/signals"
)

func main() {
	stopCh := signals.SetupSignalHandler()
	<-stopCh
	<-stopCh
}
