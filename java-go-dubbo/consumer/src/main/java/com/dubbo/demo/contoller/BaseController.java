package com.dubbo.demo.contoller;

import com.alibaba.dubbo.config.annotation.Reference;
import com.dubbo.demo.model.RequestData;
import com.dubbo.demo.model.RequestListData;
import com.dubbo.demo.model.ResponseData;
import com.dubbo.demo.model.ResponseListData;
import com.dubbo.demo.service.DataService;
//import org.apache.dubbo.config.annotation.DubboReference;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping(value="/consumer")
public class BaseController {
//http://127.0.0.1:9999/consumer/index/java
//https://blog.csdn.net/freewebsys/article/details/82780264

    //@DubboReference(version = "1.0.0",timeout = 1000, loadbalance = "roundrobin")
    //@DubboReference(version = "*",timeout = 1000)
    @Reference(version = "1.0.0",timeout = 1000, loadbalance = "roundrobin")
    private DataService dataService;

    @RequestMapping(value = "/index/{name}", method = RequestMethod.GET)
    String hello(@PathVariable("name") String name) {
        RequestData req = new RequestData();
        req.setTable(name);
        ResponseData responseData = dataService.DataList(req);
        System.out.println(responseData.getTable());
       /* Map<String,String>[] data = responseData.getData();
        if (data != null) {
            for (int i = 0; i < data.length; i++) {
                System.out.println(data[i]);
            }
        }*/
        return "Hello World " + responseData+ "!";
    }

    @RequestMapping(value = "/index2/{name}", method = RequestMethod.GET)
    String hello2(@PathVariable("name") String name) {

        RequestListData req = new RequestListData();
        req.setRequestData(new RequestData(name));
        List<Long> list = new ArrayList<Long>();
        list.add(11L);
        list.add(12L);
        list.add(100L);
        req.setListId(list);
        ResponseListData responseData = dataService.DataListMethod(req);
        System.out.println(responseData.getTable());
        return "Hello World " + responseData+ "!";
    }

}
