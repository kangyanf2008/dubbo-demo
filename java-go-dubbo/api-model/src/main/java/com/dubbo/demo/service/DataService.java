package com.dubbo.demo.service;

import com.dubbo.demo.model.RequestData;
import com.dubbo.demo.model.RequestListData;
import com.dubbo.demo.model.ResponseData;
import com.dubbo.demo.model.ResponseListData;

/**
 * 数据接口
 */
public interface DataService {
    public ResponseData DataList(RequestData data);
    public ResponseListData DataListMethod(RequestListData data);
}
