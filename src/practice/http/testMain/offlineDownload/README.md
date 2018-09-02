# 研究go的下载器

## 安装
* 依赖fasthttp，下载`go get -u github.com/valyala/fasthttp`

### 运行
* `go run src/practice/http/testMain/offlineDownload/main.go --host="0.0.0.0" --port="8081"`

## 功能流程分析
### 路由
* 通过`ctx.Path()`的值进行判断，其中`ctx`的类型是`*fasthttp.RequestCtx`


## 参考资料
* https://github.com/ilanyu/offLineDownloader
