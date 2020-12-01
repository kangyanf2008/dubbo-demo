package com.dubbo.demo.bootstrap;

//import org.apache.dubbo.config.spring.context.annotation.DubboComponentScan;
//import org.apache.dubbo.config.spring.context.annotation.EnableDubboConfig;
import com.alibaba.dubbo.config.spring.context.annotation.DubboComponentScan;
import com.alibaba.dubbo.config.spring.context.annotation.EnableDubbo;
import com.alibaba.dubbo.config.spring.context.annotation.EnableDubboConfig;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

import java.io.IOException;

@SpringBootApplication
@EnableDubboConfig
@DubboComponentScan({"com.dubbo.demo.service"})
@ComponentScan(value = {"com.dubbo.demo.contoller"})
public class ConsumerBootstrap {
    public static void main(String[] args) throws IOException {
        SpringApplication.run(ConsumerBootstrap.class, args);
        System.in.read();
    }
}
