package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	clientset "github.com/monkeyboy123/custom-controller/pkg/client/clientset/versioned"
	informers "github.com/monkeyboy123/custom-controller/pkg/client/informers/externalversions"
	"github.com/monkeyboy123/custom-controller/pkg/signals"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL   string
	kuberconfig string
)

func main() {
	flag.Parse()
	stopCh := signals.SetupSignalHandler()

	glog.Info("gogoggoog flag")
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubeClient clientset: %s", err.Error())
	}
	networkClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}
	networkInformerFactory := informers.NewSharedInformerFactory(networkClient, time.Second*30)

	glog.Info("gogoggoog networkClient")
	controller := NewController(kubeClient, networkClient, networkInformerFactory.Samplecrd().V1().Networks())
	go networkInformerFactory.Start(stopCh)
	if err = controller.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}

}

func init() {
	flag.StringVar(&kuberconfig, "kubeconfig", "", "path to kubecongfig. only require if out-of-cluster")
	flag.StringVar(&masterURL, "master", "", "the address of the kubernetes API server. Override any value in kubeconfig . Only required if out-of-cluster")
}
