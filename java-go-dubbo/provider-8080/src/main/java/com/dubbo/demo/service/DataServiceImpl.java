package com.dubbo.demo.service;

import com.dubbo.demo.model.RequestData;
import com.dubbo.demo.model.ResponseData;
import org.apache.dubbo.config.annotation.DubboService;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@DubboService(version = "1.0.0")
//@DubboService(version = "*")
@Service
@Component
public class DataServiceImpl implements DataService {

    public ResponseData DataList(RequestData req) {
        System.out.println("call "+req.getTable());
        ResponseData responseData = new ResponseData();
        responseData.setTable("java provider "+req.getTable());

        Map<String,String> d1 = new HashMap<String, String>();
        d1.put("id","1");
        d1.put("userName","java 张三");

        Map<String,String> d2 = new HashMap<String, String>();
        d2.put("id","2");
        d2.put("userName","java 李四");

        Map<String, String>[] data = new HashMap[2];
        data[0] = d1;
        data[1] = d2;
        responseData.setData(d1);
        System.out.println("end call "+req.getTable());
        return responseData;
    }

}
