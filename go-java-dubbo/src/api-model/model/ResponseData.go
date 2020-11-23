package model

import hessian "github.com/apache/dubbo-go-hessian2"

func init()  {
	hessian.RegisterPOJO(&ResponseData{})
}

type ResponseData struct {
	Table string `hessian:"table"`
	//Data []map[string]string
	Data map[string]string `hessian:"data"`
}

func (ResponseData) JavaClassName() string {
	return "com.dubbo.demo.model.ResponseData"
}
