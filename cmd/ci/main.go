package main

import (
	"flag"

	"github.com/laik/yce-cloud-extensions/pkg/common"
	"github.com/laik/yce-cloud-extensions/pkg/configure"
	ctl "github.com/laik/yce-cloud-extensions/pkg/controller"
	"github.com/laik/yce-cloud-extensions/pkg/datasource/k8s"
	fe "github.com/laik/yce-cloud-extensions/pkg/utils/file"
)

func needInit() (*configure.InstallConfigure, error) {
	var k8sJSONData []byte
	if common.InCluster {
		configure.SetTheAppRuntimeMode(configure.InCluster)
	} else {
		d, err := fe.NewIConvert(common.DefaultConfigFile).Convert()
		if err != nil {
			return nil, err
		}
		k8sJSONData = d
	}
	cfg, err := configure.NewInstallConfigure(k8s.NewResources(nil), k8sJSONData)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func main() {
	flag.Parse()
	cfg, err := needInit()
	if err != nil {
		panic(err)
	}

	srv := ctl.NewCIController(cfg)
	stop := make(chan struct{})
	if err := srv.Run(":8080", stop); err != nil {
		panic(err)
	}
}
