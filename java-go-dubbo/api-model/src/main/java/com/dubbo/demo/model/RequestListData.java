package com.dubbo.demo.model;

import lombok.Data;

import java.io.Serializable;
import java.util.List;

@Data
public class RequestListData implements Serializable {
    public List<Long> listId;
    public RequestData requestData;
}
