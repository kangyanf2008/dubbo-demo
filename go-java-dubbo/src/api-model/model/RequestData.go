package model

import hessian "github.com/apache/dubbo-go-hessian2"

func init()  {
	hessian.RegisterPOJO(&RequestData{})
}
type RequestData struct {
	Table string `hessian:"table"`
}
func (RequestData) JavaClassName() string {
	return "com.dubbo.demo.model.RequestData"
}