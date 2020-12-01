package com.dubbo.demo.bootstrap;

//import org.apache.dubbo.config.spring.context.annotation.DubboComponentScan;
//import org.apache.dubbo.config.spring.context.annotation.EnableDubboConfig;
import com.alibaba.dubbo.config.spring.context.annotation.DubboComponentScan;
import com.alibaba.dubbo.config.spring.context.annotation.EnableDubboConfig;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.io.IOException;

@SpringBootApplication
@EnableDubboConfig
@DubboComponentScan({"com.dubbo.demo.service"})
public class ProviderBootstrap {
    public static void main(String[] args) throws IOException {
        SpringApplication.run(ProviderBootstrap.class, args);
        System.in.read();
    }
}
