package com.dubbo.demo.service;

import com.dubbo.demo.model.RequestData;
import com.dubbo.demo.model.ResponseData;

/**
 * 数据接口
 */
public interface DataService {
    public ResponseData DataList(RequestData data);
}
