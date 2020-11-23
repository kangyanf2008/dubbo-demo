package service

import (
	"context"
	"fmt"
	"go-java-dubbo/api-model/model"
)

type DataServiceProvider struct {
}


func (u *DataServiceProvider) DataList(ctx context.Context, reqData *model.RequestData) (*model.ResponseData, error) {
	fmt.Println("call:%#v \n", reqData)
	responseData := &model.ResponseData{}
	responseData.Table = "go provider " + reqData.Table

	data := make([]map[string]string, 2)

	d1 := make(map[string]string)
	d1["id"]="3"
	d1["userName"]="go 张三"

	d2 := make(map[string]string)
	d2["id"]="4"
	d2["userName"]="go 李四"

	data[0] = d1
	data[1] = d2
	//responseData.Data = data
	responseData.Data = d1

	fmt.Printf("rsp:%#v \n", responseData)
	return responseData, nil
}

func (u *DataServiceProvider) Reference() string {
	return "com.dubbo.demo.service.DataService"
}
