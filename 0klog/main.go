package main

import (
	"flag"
	"k8s.io/klog/v2"
)

func main() {
	// Initialize klog
	klog.InitFlags(nil)
	defer klog.Flush()

	// Parse command line flags
	flag.Set("v", "2") // Set the verbosity level to 2
	flag.Parse()

	klog.Info("This is an info message")
	klog.Warning("This is a warning message")
	klog.Error("This is an error message")

	// Verbose logging
	klog.V(1).Info("This is a level 1 verbosity message")
	klog.V(2).Info("This is a level 2 verbosity message")
	klog.V(3).Info("This is a level 3 verbosity message")
}
