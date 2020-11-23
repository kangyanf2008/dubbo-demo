package main

import (
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/apache/dubbo-go/common/logger"
	"go-java-dubbo/api-model/model"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/dubbo"
	_ "github.com/apache/dubbo-go/registry/protocol"

	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
	"go-java-dubbo/consumer/controller"
)

var (
	survivalTimeout = int(3e9)
	server *http.Server
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hellword"))
}

// 		export CONF_CONSUMER_FILE_PATH="xxx"
// 		export APP_LOG_CONF_FILE="xxx"
func main() {

/*	config.SetConsumerService(DataService)

 */
	hessian.RegisterPOJO(&model.RequestData{})
	hessian.RegisterPOJO(&model.ResponseData{})
	config.Load()
	time.Sleep(3e9)

/*	for i:=0; i<5; i++ {
		responseData, err := service.DataService.DataList(context.TODO(), &model.RequestData{Table:"go"})
		if err != nil {
			panic(err)
		}
		fmt.Println(responseData)
	}
*/
	controller.Run()

	//这里实现了远程获取pprof数据的接口
	go func() {
		/*
		pprof.HandleFunc("/debug/pprof/", pprof.Index)
		pprof.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		pprof.HandleFunc("/debug/pprof/profile", pprof.Profile)
		pprof.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		pprof.HandleFunc("/debug/pprof/trace", pprof.Trace)
		*/
		fmt.Print("==================",http.ListenAndServe("0.0.0.0:7777",nil))
	}()

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
			controller.Shutdown()
			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider app exit now...")
			return
		}
	}
}

