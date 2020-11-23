package com.dubbo.demo.model;

import lombok.Data;

import java.io.Serializable;

@Data
public class RequestData implements Serializable {
    public String table;
}
