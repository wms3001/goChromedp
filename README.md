# goChromedp
## 简介
Chromedp简单使用，实现截图和打印pdf
## 使用
```
go get github.com/wms3001/goChromedp
```
## 实例
1. 截图
```go
goChr := &GoChromedp{}
goChr.Url = "https://www.oschina.net"
goChr.Quality = 100
ctx, cancel := goChr.Ctx()
resp := goChr.Screenshot(ctx, cancel)
res, _ := json.Marshal(resp)
log.Println(string(res))
```
2. pdf
```go
goChr := &GoChromedp{}
goChr.Url = "https://www.oschina.net"
ctx, cancel := goChr.Ctx()
resp := goChr.PrintPdf(ctx, cancel)
res, _ := json.Marshal(resp)
log.Println(string(res))
```
