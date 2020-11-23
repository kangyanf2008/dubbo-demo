package service

import (
	"context"
	"github.com/apache/dubbo-go/config"
	"go-java-dubbo/api-model/model"
)

var DataService = new(DataServiceProvider)

func init() {
	config.SetConsumerService(DataService)
}

type DataServiceProvider struct {
	//DataList func(req *model.RequestData) (*model.ResponseData, error)
	DataList func(ctx context.Context, req *model.RequestData, rsp *model.ResponseData) error `dubbo:"DataList"`
}

func (u *DataServiceProvider) Reference() string {
	return "com.dubbo.demo.service.DataService"
}