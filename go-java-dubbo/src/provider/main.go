package main

import (
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/apache/dubbo-go/common/logger"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/protocol/dubbo"
	_ "github.com/apache/dubbo-go/registry/protocol"
	"go-java-dubbo/api-model/model"
	"go-java-dubbo/provider/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	_ "github.com/apache/dubbo-go/filter/filter_impl"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
)

var (
	survivalTimeout = int(3e9)
)

// 		export CONF_PROVIDER_FILE_PATH="xxx"
// 		export APP_LOG_CONF_FILE="xxx"
func main() {
	config.SetProviderService(new(service.DataServiceProvider))
	hessian.RegisterPOJO(&model.RequestData{})
	hessian.RegisterPOJO(&model.ResponseData{})
	config.Load()

	initSignal()
}

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})
			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider app exit now...")
			return
		}
	}
}