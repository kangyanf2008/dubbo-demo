package com.dubbo.demo.model;

import lombok.Data;

import java.io.Serializable;
import java.util.Map;

@Data
public class ResponseData implements Serializable {
    private String table;
    //private Map<String,String>[] data;
    private Map<String,String> data;
}
