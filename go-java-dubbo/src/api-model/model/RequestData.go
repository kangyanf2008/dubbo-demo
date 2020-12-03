package model

import hessian "github.com/apache/dubbo-go-hessian2"

func init()  {
	hessian.RegisterPOJO(&RequestListData{})
	hessian.RegisterPOJO(&RequestData{})
}
type RequestData struct {
	Table string `hessian:"table"`
}
func (RequestData) JavaClassName() string {
	return "com.dubbo.demo.model.RequestData"
}

type RequestListData struct {
	ListId []int64 `hessian:"listId"`
	RequestData *RequestData `hessian:"requestData"`
}
func (RequestListData) JavaClassName() string {
	return "com.dubbo.demo.model.RequestListData"
}
