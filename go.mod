module github.com/liyanbing/golang-pack

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.55.0
	go.opencensus.io => github.com/census-instrumentation/opencensus-go v0.22.3
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200323165209-0ec3e9974c59
	golang.org/x/lint => github.com/golang/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/mod => github.com/golang/mod v0.2.0
	golang.org/x/net => github.com/golang/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync => github.com/golang/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200331124033-c3d80250170d
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20191024005414-555d28b269f0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20200331202046-9d5940d49312
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191204190536-9bdfabe68543
	google.golang.org/appengine => github.com/allenlulu/appengine v0.0.0-20171126154002-d35ccace8f27
	google.golang.org/genproto => github.com/googleapis/go-genproto v0.0.0-20200331122359-1ee6d9798940
	google.golang.org/grpc => github.com/grpc/grpc-go v1.28.0
)

require (
	github.com/go-kit/kit v0.10.0
	github.com/hashicorp/consul/api v1.4.0
	github.com/opentracing/opentracing-go v1.1.0
	google.golang.org/grpc v1.28.0
)
