package com.dubbo.demo.model;

import lombok.Data;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

@Data
public class ResponseListData implements Serializable {
    private String table;
    private List<Long> data;
}
