module go-java-dubbo

require (
	github.com/RoaringBitmap/roaring v0.4.23 // indirect
	github.com/apache/dubbo-go v1.5.4
	github.com/apache/dubbo-go-hessian2 v1.7.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dubbogo/gost v1.9.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5 // indirect
	github.com/openzipkin/zipkin-go v0.2.2 // indirect
	github.com/uber/jaeger-client-go v2.22.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
)

replace (
	github.com/envoyproxy/go-control-plane => github.com/envoyproxy/go-control-plane v0.8.0
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20140225173054-eb6ee6f84d0a
)

go 1.13
