package service

import "go-java-dubbo/api-model/model"

type DataService interface {
	DataList(ctx context.Context, data *model.RequestData) (*model.ResponseData, error)
	//DataList(data *model.RequestData) (*model.ResponseData, error)
}
