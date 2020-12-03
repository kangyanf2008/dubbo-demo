package service

import (
	"context"
	"go-java-dubbo/api-model/model"
)

type DataService interface {
	DataList(ctx context.Context, data *model.RequestData) (*model.ResponseData, error)
	DataListMethod(ctx context.Context, data *model.ResponseListData) (*model.ResponseListData, error)
}
