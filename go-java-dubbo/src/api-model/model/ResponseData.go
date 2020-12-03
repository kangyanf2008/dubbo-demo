package model

import hessian "github.com/apache/dubbo-go-hessian2"

func init()  {
	hessian.RegisterPOJO(&ResponseData{})
	hessian.RegisterPOJO(&ResponseListData{})
}

type ResponseData struct {
	Table string `hessian:"table"`
	//Data []map[string]string
	Data map[string]string `hessian:"data"`
}

func (ResponseData) JavaClassName() string {
	return "com.dubbo.demo.model.ResponseData"
}

type ResponseListData struct {
	Table string  `hessian:"table"`
	Data  []int64 `hessian:"data"`
}
func (ResponseListData) JavaClassName() string {
	return "com.dubbo.demo.model.ResponseListData"
}